package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Post struct {
	Title string `json:"title"`
}

func main() {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Ошибка HTTP: %s", resp.Status)
	}

	var posts []Post

	err = json.NewDecoder(resp.Body).Decode(&posts)
	if err != nil {
		log.Fatal(err)
	}
	

	fmt.Println("Всего постов:", len(posts))

	if len(posts) > 0 {
		fmt.Println("Первый пост:", posts[0].Title)
		fmt.Println("Последний пост:", posts[len(posts)-1].Title)
	}
}
