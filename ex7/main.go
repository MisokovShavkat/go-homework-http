package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	req, err := http.NewRequest(
		http.MethodGet,
		"https://httpbin.org/headers",
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("X-Student", "Shavkat")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))

	if strings.Contains(string(body), `"X-Student"`) {
		fmt.Println("Заголовок найден!")
	}
}
