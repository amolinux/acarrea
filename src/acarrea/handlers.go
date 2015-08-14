package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
    "database/sql"
	"log"
//	"io"
//	"io/ioutil"
)

const (
    DB_HOST = "tcp(localhost:3306)"
    DB_NAME = "dbacarrea"
    DB_USER = /*"root"*/ "acarrea"
    DB_PASS = /*""*/ "acarrea123"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to Acarrea!\n")
}

func ClienteIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	
	var_clients := Clients{}
	
	dsn := DB_USER + ":" + DB_PASS + "@" + DB_HOST + "/" + DB_NAME + "?charset=utf8"
	db, _ := sql.Open("mysql", dsn)
	defer db.Close()

	// Connect and check the server version
	rows, err := db.Query("select id, name, last_name, address, phone, email, active, id_document, ac_document_type_id from ac_client")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		c := Client{}
		err := rows.Scan(&c.Id, &c.Name, &c.Last_name, &c.Address, &c.Phone, &c.Email, &c.Active, &c.Id_document, &c.Document_type_id)
		if err != nil {
			log.Fatal(err)
		}
		var_clients = append(var_clients, c)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    
	if err := json.NewEncoder(w).Encode(var_clients); err != nil {
		panic(err)
	}
}

func ClienteShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var clientId int
	var err error
	if clientId, err = strconv.Atoi(vars["clientId"]); err != nil {
		panic(err)
	}
	client := RepoFindClient(clientId)
	if client.Id > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(client); err != nil {
			panic(err)
		}
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)

}
