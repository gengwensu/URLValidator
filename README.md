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


    ```
    $ curl http://localhost:8081/urlVal/malwareType?url="test1.com:8000/this/is/a/test"
    {"hostname":"test1.com","type":"malware"}

    $ curl http://localhost:8081/urlVal/malwareType?url="test2.com:8000/this/is/a/test"
    {"hostname":"test2.com","type":"clean"}
    
    ```
 
the server should respond with 404 to all other requests not listed above
 
 # environment & build
 require Go and mySql
 dbinit.sql should be run after the installation of mysql to set up the db environment
 
$ go build ../src/github.com/gengwensu/URLValidator/urlVal.go ../src/github.com/gengwensu/URLValidato r/queryDB.go

$./urlVal &
...

# Unit test
require sqlmock; to install
$ go get gopkg.in/DATA-DOG/go-sqlmock.v1

$ go test                                                     
Replacing cache entry test9.com cache count 1 with test2.com. 
Replacing cache entry test9.com cache count 1 with test6.com. 
PASS                                                          
ok      github.com/gengwensu/URLValidator       0.043s        

# Implementation
A cache with MAXCACHEENTRY entries are added to improve the performance. urlVal will check the cache before querying the malware table in the database. If the request hostname is not in the cache, the result will be cached when there're still spaces in the cache. Otherwise, the least used entry will be replaced with the result.
