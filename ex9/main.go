package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	codes := []int{200, 201, 404, 500, 503}

	for _, code := range codes {
		url := fmt.Sprintf("https://httpbin.org/status/%d", code)

		resp, err := client.Get(url)
		if err != nil {
			fmt.Printf("%d: ошибка клиента\n", code)
			continue
		}

		resp.Body.Close()

		switch {
		case resp.StatusCode >= 200 && resp.StatusCode < 300:
			fmt.Printf("%d: успех\n", code)
		case resp.StatusCode >= 400 && resp.StatusCode < 500:
			fmt.Printf("%d: ошибка клиента\n", code)
		case resp.StatusCode >= 500 && resp.StatusCode < 600:
			fmt.Printf("%d: ошибка сервера\n", code)
		}
	}
}
