package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// ErrUserNotFound User Not Found Error
// Action layer
var ErrUserNotFound = errors.New("User not found")

type User struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type MemoryUserStorage struct {
	store map[string]*User
}

// UserStorer any type with these methods are also type of UserStorer
type UserStorer interface {
	Get(ctx context.Context, email string) (*User, error)
	Save(ctx context.Context, user *User) error
}

func NewMemoryUserStorage() *MemoryUserStorage {
	return &MemoryUserStorage{
		store: map[string]*User{},
	}
}

func (ms *MemoryUserStorage) Get(ctx context.Context, email string) (*User, error) {
	if u, ok := ms.store[email]; ok {
		return u, nil
	}
	return nil, ErrUserNotFound
}

func (ms *MemoryUserStorage) Save(ctx context.Context, user *User) error {
	ms.store[user.Email] = user
	return nil
}

// Buisness Layer
type RegisterParams struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (rp *RegisterParams) Validate() error {
	if rp.Email == "" {
		return errors.New("Emails cannot be empty")
	}

	if !strings.ContainsRune(rp.Email, '@') {
		return errors.New("Email is in incorrect format @ not included")
	}

	if rp.Name == "" {
		return errors.New("Name cannot be empty")
	}

	return nil
}

type UserService interface {
	Register(ctx context.Context, params *RegisterParams) error
	GetByEmail(ctx context.Context, email string) (*User, error)
}

var ErrEmailExists = errors.New("Email is already in use.")

type UserServiceImpl struct {
	userStorage UserStorer
}

func NewUserServiceImpl(userStr UserStorer) *UserServiceImpl {
	return &UserServiceImpl{
		userStorage: userStr,
	}
}

func (us *UserServiceImpl) Register(ctx context.Context, params *RegisterParams) error {
	_, err := us.userStorage.Get(ctx, params.Email)
	if err == nil {
		return ErrEmailExists
	} else if err != ErrUserNotFound {
		return err
	}
	return us.userStorage.Save(ctx, &User{
		Email: params.Email,
		Name:  params.Name,
	})
}

func (us *UserServiceImpl) GetByEmail(ctx context.Context, email string) (*User, error) {
	return us.userStorage.Get(ctx, email)
}

// Access Layer
type JSONOverHTTP struct {
	router      *http.ServeMux
	userService UserService
}

func NewJSONOverHTTP(userService UserService) *JSONOverHTTP {
	r := http.NewServeMux()
	jsonOverHTTP := &JSONOverHTTP{
		router:      r,
		userService: userService,
	}

	r.HandleFunc("/register", jsonOverHTTP.Register)
	r.HandleFunc("/user", jsonOverHTTP.GetUser)
	return jsonOverHTTP
}

// To implement HandlerFunc
func (jsonOverHTTP *JSONOverHTTP) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	jsonOverHTTP.router.ServeHTTP(w, r)
}

func (jsonOverHTTP *JSONOverHTTP) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Register requires a POST request", http.StatusMethodNotAllowed)
		return
	}

	params := &RegisterParams{}

	err := json.NewDecoder(r.Body).Decode(params)
	if err != nil {
		http.Error(w, "Enable to read your request", http.StatusBadRequest)
		return
	}

	err = params.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = jsonOverHTTP.userService.Register(r.Context(), params)
	if err == ErrEmailExists {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}

func (jsonOverHTTP *JSONOverHTTP) ValidateEmail(email string) error {
	if email == "" {
		return errors.New("Email cannot be empty")
	}

	if !strings.ContainsRune(email, '@') {
		return errors.New("Email doesnt contain @")
	}

	return nil
}

func (jsonOverHTTP *JSONOverHTTP) GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "GetUser requires a GET METHOD", http.StatusMethodNotAllowed)
		return
	}

	email := r.FormValue("email")
	err := jsonOverHTTP.ValidateEmail(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := jsonOverHTTP.userService.GetByEmail(r.Context(), email)
	if err == ErrUserNotFound {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	userStr := NewMemoryUserStorage()
	userService := NewUserServiceImpl(userStr)
	jsonOverHTTP := NewJSONOverHTTP(userService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	err := http.ListenAndServe(":"+port, jsonOverHTTP)
	if err != nil {
		fmt.Println("Error Serving HTTP Server:", err)
	}
}
