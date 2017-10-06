# URLValidator
A simple URL Malware Validator with limited API 

# Usecase: 
a web application using a HTTP proxy to scan traffic looking for malware URLs. Before allowing HTTP connections to be made, this proxy asks urlVal that maintains several databases of malware URLs if the resource being requested is known to contain malware.

# API:
The urlValidator will run on http://localhost:8081 and will support the following REST APIs:
1. GET /urlVal/
    returns "url Validation service"

2. GET /urlVal/validate/{hostname_and_port}/{original_path_and_query_string}
    returns a JSON string of "clean" or "{malware type}" - "malware" initially

    example:
    =======
    ```
    curl http://localhost:8081/urlVal/validate/{hostname_and_port}/{original_path_and_query_string}

    {"url": {hostname}, "type": "clean"} or {"url": {hostname}, "type": "malware"}
    ```
 
 
the server should respond with 404 to all other requests not listed above
 
 # environment & installation
 require Go
 go build ...
