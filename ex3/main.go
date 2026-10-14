package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	Title string `json:"title"`
}

func main() {
	urls := []string{
		"https://jsonplaceholder.typicode.com/todos/5",
		"https://jsonplaceholder.typicode.com/todos/99999",
	}

	for _, url := range urls {
		resp, err := http.Get(url)

		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}

		fmt.Println("Статус-код:", resp.StatusCode)

		if resp.StatusCode == http.StatusOK {
			var todo Todo
			err := json.NewDecoder(resp.Body).Decode(&todo)
			if err != nil {
				fmt.Println("Ошибка JSON:", err)
				resp.Body.Close()
				continue
			}
			fmt.Println("Title:", todo.Title)
		} else if resp.StatusCode == http.StatusNotFound {
			fmt.Println("Задача не найдена")
		}

		resp.Body.Close()
	}
}
