<h1>Curl commands for GET and POST</h1>
<h2>Flags:</h2>
<ul>
<li>-v => verbose flag for more information on sent and received</li>
<li>-XGET => default curl command is get</li>
<li>-XPOST => for post method use with -d to send some data in body or through query "localhost:3000/user?email=value"(always but URL in string if there is some query parameters)</li>
<li>-d => -d '{ "jsonKey": "jsonValue" }' used to send some json data </li>
<li>-H => Header Content-type Example: "Content-type: application/json"</li>
</ul>
