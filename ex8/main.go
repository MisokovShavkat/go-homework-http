package main

import (
	"log"
	"net/http"
	"time"
)

func main() {

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get("https://httpbin.org/delay/2")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	log.Printf("Статус ответа: %s", resp.Status)
}
