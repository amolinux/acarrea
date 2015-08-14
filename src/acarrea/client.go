package main

type Client struct {
	Id int `json:"id"`
    Name string `json:"name"`
    Last_name string `json:"last_name"`
    Address string `json:"address"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Active string `json:"active"`
	Id_document string	`json:"id_document"`
	Document_type_id int `json:"document_type_id"`
}

type Clients []Client

