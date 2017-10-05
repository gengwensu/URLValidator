# URLValidator
urlVal is a URL Malware Validator 
# Usecase: 
a web application using a HTTP proxy to scan traffic looking for malware URLs. Before allowing HTTP connections to be made, this proxy asks urlVal that maintains several databases of malware URLs if the resource being requested is known to contain malware.

# API:
 GET /urlinfo/1/{hostname_and_port}/{original_path_and_query_string}
 return a JSON string of "clean" or "{malware type}" - "malware" initially
 
 # test case:
 curl ...
 
 # environment & installation
 require Go
 go build ...
