package main

import "fmt"

var currentId int

var clients Clients

// Give us some seed data
func init() {
	RepoCreateClient(Client{Name: "Write presentation"})
	RepoCreateClient(Client{Name: "Host meetup"})
}

func RepoFindClient(id int) Client {
	for _, t := range clients {
		if t.Id == id {
			return t
		}
	}
	// return empty Todo if not found
	return Client{}
}

//this is bad, I don't think it passes race condtions
func RepoCreateClient(t Client) Client {
	currentId += 1
	t.Id = currentId
	clients = append(clients, t)
	return t
}

func RepoDestroyClient(id int) error {
	for i, t := range clients {
		if t.Id == id {
			clients = append(clients[:i], clients[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}