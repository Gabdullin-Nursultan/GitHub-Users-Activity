package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type GitHubEvents struct {
	Type string
	Repo struct {
		Name string
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Введите имя пользователя:")
		return
	}

	username := os.Args[1]
	url := fmt.Sprintf("https://api.github.com/users/%v/events", username)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка получения запроса: ", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения записи файла", err)
		return
	}

	var events []GitHubEvents
	err = json.Unmarshal(body, &events)
	if err != nil {
		fmt.Println("Ошибка парсинга JSON", err)
		return
	}

	for _, event := range events {
		fmt.Printf("%v в %v\n", event.Type, event.Repo.Name)
	}
}
