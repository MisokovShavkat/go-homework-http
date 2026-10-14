package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type User struct {
	Name    string  `json:"name"`
	Company Company `json:"company"`
}

type Company struct {
	Name string `json:"name"`
}

func main() {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest(
		http.MethodGet,
		"https://jsonplaceholder.typicode.com/users",
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var users []User

	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		log.Println(
			"Name:", user.Name,
			"Company:", user.Company.Name,
		)
	}
}
