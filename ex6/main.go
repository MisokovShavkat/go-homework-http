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

	url := "https://jsonplaceholder.typicode.com/posts?userId=3"

	resp, err := client.Get(url)
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

	fmt.Println("Количество постов:", len(posts))

	for i, post := range posts {
		fmt.Printf("%d. %s\n", i+1, post.Title)
	}
}
