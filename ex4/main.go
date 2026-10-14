package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Todo struct {
	Completed bool `json:"completed"`
}

func main() {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Ошибка HTTP: %s", resp.Status)
	}

	var todos []Todo

	err = json.NewDecoder(resp.Body).Decode(&todos)
	if err != nil {
		log.Fatal(err)
	}

	completed := 0
	notCompleted := 0

	for _, todo := range todos {
		if todo.Completed {
			completed++
		} else {
			notCompleted++
		}
	}

	fmt.Println("Выполнено:", completed)
	fmt.Println("Не выполнено:", notCompleted)
	fmt.Println("Всего задач:", len(todos))
}
