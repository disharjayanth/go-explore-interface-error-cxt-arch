package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func makeRequest(ctx context.Context, route string) (string, error) {
	deadlineTime, ok := ctx.Deadline()
	if ok && time.Until(deadlineTime) < 100*time.Millisecond {
		return "", fmt.Errorf("Deadline too near")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, route, nil)
	if err != nil {
		return "", fmt.Errorf("Cannot create context request: ", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("Cannot make a context request to google.com", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Bad status code: %d", res.StatusCode)
	}

	sb, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("Error reading from response body: ", err)
	}

	return string(sb), nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 320*time.Millisecond)
	defer cancel()

	res, err := makeRequest(ctx, "https://www.google.com")
	if err != nil {
		panic(err)
	}

	fmt.Println("Response: ", res)
}
