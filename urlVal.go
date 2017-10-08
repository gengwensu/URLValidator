/*
urlVal: a malware validator service that accept a HTTP GET request with an embbedded url.
It will check the hostname of the embbedded url against a malware db and return the malware
type for the url.
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var malDb = []string{"test1.com", "196.132.1.1"}

//AppContext contains all global variables that are shared among packages
type AppContext struct {
	MalMap     map[string]string // map with hostname as key and malware type as value
	CacheCount map[string]int    // counts for each cache entry
	DbHandler  *sql.DB           // db handle to mySql
}

const MAXCACHEENTRY int = 2

type Response struct {
	Hostname    string `json:"hostname"`
	MalwareType string `json:"type"`
}

func main() {
	aMap := map[string]string{}
	cMap := map[string]int{}
	for i, s := range malDb {
		if i < MAXCACHEENTRY {
			aMap[s] = "malware"
			cMap[s] = 1
		}
	}
	//start http server
	dbh, err := sql.Open("mysql", "root@/maldb")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer dbh.Close()

	globalVar := AppContext{
		MalMap:     aMap,
		DbHandler:  dbh,
		CacheCount: cMap,
	}

	log.Fatal(http.ListenAndServe("localhost:8081", &globalVar))
}

func (ds *AppContext) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var hostname string
	switch req.URL.Path {
	case "/urlVal", "/urlVal/":
		if req.Method == "GET" {
			fmt.Fprint(w, "url Validation service\n") // return signature of the service
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}

	case "/urlVal/malwareType":
		if req.Method == "GET" {
			urlString := req.URL.Query().Get("url") //get the embbedded url
			// fmt.Fprintln(w, urlString)
			hostString := strings.SplitN(urlString, "/", 2) //split into two parts
			// fmt.Fprintln(w, hostString[0])
			//the first part is the hostname:port
			if hostString[0] != "" {
				// get client hostname
				var err error
				hostname, _, err = net.SplitHostPort(hostString[0])
				if err != nil {
					fmt.Fprintf(w, "hostname:port %v is not valide\n", urlString)
				} else {
					resp := &Response{}
					//see if hostname is in the malware list
					resp.MalwareType = ds.QueryDB(hostname)
					resp.Hostname = hostname
					rout, _ := json.Marshal(resp)
					fmt.Fprintf(w, string(rout))
				}
			} else {
				fmt.Fprintf(w, "url %s missing hostname:port\n", req.URL)
			}
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}

	default:
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "http 404, %s invalid. Only /urlVal/validate is allowed.\n", req.URL)
	}
}
