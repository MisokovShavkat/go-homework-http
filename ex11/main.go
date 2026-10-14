package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	req, err := http.NewRequest(
		http.MethodDelete,
		"https://jsonplaceholder.typicode.com/posts/1",
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

	fmt.Println("Status code: ", resp.StatusCode)
}
