package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Printf(todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

// contoh ga pakai pointer
func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("invalid input")
	}

	return Todo{
		Text: content,
	}, nil
}

// package todo

// import (
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// 	"os"
// )

// type Todo string

// func (todo Todo) Display() {
// 	fmt.Println(string(todo))
// }

// func (todo Todo) Save() error {
// 	fileName := "todo.json"

// 	jsonData, err := json.Marshal(todo)
// 	if err != nil {
// 		return err
// 	}

// 	return os.WriteFile(fileName, jsonData, 0644)
// }

// func New(content string) (Todo, error) {
// 	if content == "" {
// 		return "", errors.New("invalid input")
// 	}
// 	return Todo(content), nil
// }
