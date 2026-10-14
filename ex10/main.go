package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"net/http"

	"time"
)

type Post struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"userId"`
}

type PostResponse struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"userId"`
}

func main() {

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	post := Post{
		Title:  "My first post",
		Body:   "Text my post",
		UserID: 1,
	}
	jsonData, err := json.Marshal(post)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Post(
		"https://jsonplaceholder.typicode.com/posts",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var result PostResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Статус-код:", resp.StatusCode)
	fmt.Println("ID созданного поста:", result.ID)

}
