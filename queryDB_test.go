package main

import (
	"log"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

var tAC AppContext

func initTestEnv() {
	am := map[string]string{}
	cm := map[string]int{}
	am["test1.com"] = "malware"
	cm["test1.com"] = 2
	am["test9.com"] = "clean"
	cm["test9.com"] = 1

	tAC = AppContext{
		MalMap:     am,
		DbHandler:  nil,
		CacheCount: cm,
	}
}

func TestQueryDBHitCache(t *testing.T) { // return result from cache
	var tests = []struct {
		iHostname string
		want      string
	}{
		{"test1.com", "malware"},
		{"test9.com", "clean"},
	}

	initTestEnv()
	for _, test := range tests {
		if result := tAC.QueryDB(test.iHostname); result != test.want {
			t.Errorf("QueryDB(%s) retruns %s; want %s\n", test.iHostname, result, test.want)
		}
	}
}

func TestQueryDBHitDB(t *testing.T) { // return result from db; replace cache
	var tests = []struct {
		iHostname string
		want      string
	}{
		{"test2.com", "phising"},
	}

	initTestEnv()
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	tAC.DbHandler = db
	columns := []string{"type"}
	mock.ExpectQuery("SELECT type FROM malware WHERE hostname = (.*)").
		WithArgs("test2.com").
		WillReturnRows(sqlmock.NewRows(columns).AddRow("phising"))

	for _, test := range tests {
		if result := tAC.QueryDB(test.iHostname); result != test.want {
			t.Errorf("QueryDB(%s) retruns %s; want %s\n", test.iHostname, result, test.want)
		}
	}
}

func TestQueryDBNoCacheNoDB(t *testing.T) { //return clean; replace cache
	var tests = []struct {
		iHostname string
		want      string
	}{
		{"test6.com", "clean"},
	}

	initTestEnv()
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	tAC.DbHandler = db

	mock.ExpectQuery("SELECT type FROM malware WHERE hostname = (.*)").
		WithArgs("test6.com").
		WillReturnError(err)

	for _, test := range tests {
		if result := tAC.QueryDB(test.iHostname); result != test.want {
			t.Errorf("QueryDB(%s) retruns %s; want %s\n", test.iHostname, result, test.want)
		}
	}
}
