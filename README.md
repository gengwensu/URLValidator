# URLValidator
A simple URL Malware Validator with limited API 

# Usecase: 
a web application using a HTTP proxy to scan traffic looking for malware URLs. Before allowing HTTP connections to be made, this proxy asks urlVal that maintains several databases of malware URLs if the resource being requested is known to contain malware.

# API:
The urlValidator will run on http://localhost:8081 and will support the following REST APIs:
1. GET /urlVal/
    returns "url maleware Validation service"

2. GET /urlVal/malewareType?url="{hostname_and_port}/{original_path_and_query_string}"
    returns a JSON string of "clean" or "{malware type}" - "malware" only initially

    example: (malwareList = {"test1.com", "196.132.1.1"})
    =======

    ```
    $ curl http://localhost:8081/urlVal/malwareType?url="test1.com:8000/this/is/a/test"
    {"hostname":"test1.com","type":"malware"}

    $ curl http://localhost:8081/urlVal/malwareType?url="test2.com:8000/this/is/a/test"
    {"hostname":"test2.com","type":"clean"}
    
    ```
 
the server should respond with 404 to all other requests not listed above
 
 # environment & build
 require Go
 
 $ go build ../src/github.com/gengwensu/URLValidator/urlVal.go
 ./urlVal &

