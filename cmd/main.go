package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

var malDb = []string{"test1.com", "196.132.1.1"}

type AppContext struct {
	MalList []string
}

func main() {
	//start http server
	globalVar := AppContext{malDb}
	log.Fatal(http.ListenAndServe("localhost:8081", &globalVar))
}

func (db *AppContext) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var hostname, port string
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "http 404, %s invalid. Only /urlVal/validate is allowed.\n", req.URL)
	case "/urlVal", "/urlVal/":
		if req.Method == "GET" {
			fmt.Fprint(w, "url Validation service\n") // return signature of the service
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}

	default:
		if req.Method != "GET" {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			break
		}
		urlSlices := strings.Split(req.URL.Path[1:], "/") //parsing url
		if urlSlices[0] == "urlVal" && urlSlices[1] == "validate" {
			if len(urlSlices) >= 3 && urlSlices[2] != "" {
				// get client hostname
				var err error
				hostname, port, err = net.SplitHostPort(urlSlices[2])
				if err != nil {
					fmt.Fprintf(w, "hostname:port %s is valide\n", urlSlices[2])
				}
				// fmt.Fprintf(w, "Hostname: %s\n", hostname)
				// fmt.Fprintf(w, "Port: %s\n", port)
			} else {
				fmt.Fprintf(w, "url %s missing hostname:port\n", req.URL)
			}
		} else {
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "http 404, %s invalid. Only /urlVal/validate is allowed.\n", req.URL)
		}
	}

	//see if hostname is in the malware list
	flagHit := false
	for _, h := range db.MalList {
		// fmt.Fprintf(w, "in loop, h %s, hostname %s\n", h, hostname)
		if hostname == h {
			flagHit = true
			fmt.Fprintf(w, "{url: %s, type: mailware}\n", hostname)
		}
	}
	if !flagHit {
		fmt.Fprintf(w, "{url: %s, type: clean}, port %s\n", hostname, port)
	}
}
