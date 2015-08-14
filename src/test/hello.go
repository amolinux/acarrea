package main

import (
    "fmt"
    "html"
//    "log"
    "net/http"
//    "github.com/gorilla/mux"
    "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const (
    DB_HOST = "tcp(172.99.67.121:3306)"
    DB_NAME = "dbacarrea"
    DB_USER = /*"root"*/ "acarrea"
    DB_PASS = /*""*/ "acarrea123"
)

func main() {
	dsn := DB_USER + ":" + DB_PASS + "@" + DB_HOST + "/" + DB_NAME + "?charset=utf8"
	db, _ := sql.Open("mysql", dsn)
	defer db.Close()

	// Connect and check the server version
	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)

	
//    router := mux.NewRouter().StrictSlash(true)
//    router.HandleFunc("/", Index)
//    log.Fatal(http.ListenAndServe(":8080", router))    
    
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    }