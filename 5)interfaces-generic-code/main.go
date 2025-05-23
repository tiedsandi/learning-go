package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type Saver interface {
	Save() error
}

type displayer interface {
	Display()
}

type outputtable interface {
	Saver
	displayer
}

// type outputtable interface {
// 	Save() error
// 	displayer
// }

func main() {
	printSomething(1)
	printSomething(1.5)
	printSomething("Hello World!")

	title, content := getNoteData()
	todoText := getTodoData()

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Todo
	err = outputData(todo)
	if err != nil {
		return
	}

	printSomething(todo)

	// Note
	err = outputData(userNote)
	if err != nil {
		return
	}
}

func printSomething(value interface{}) {
	fmt.Println(value)

	switch v := value.(type) {
	case int:
		fmt.Println("Ini integer:", v)
	case string:
		fmt.Println("Ini string:", v)
	default:
		fmt.Println("Tipe lain:", v)
	}
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

// ada cara lain yang lebih mudah
func saveData(data Saver) error {
	err := data.Save()
	fmt.Printf("Type: %T\n", data)

	if err != nil {
		switch data.(type) {
		case *note.Note:
			fmt.Println("Saving the note failed.")
		case *todo.Todo:
			fmt.Println("Saving the todo failed.")
		default:
			fmt.Println("Saving the data failed.")
		}
		return err
	}

	switch data.(type) {
	case *note.Note:
		fmt.Println("Saving the note succeeded!")
	case *todo.Todo:
		fmt.Println("Saving the todo succeeded!")
	default:
		fmt.Println("Saving the data succeeded!")
	}

	return nil
}

func getTodoData() string {
	return getUserInput("Todo text:")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSpace(text)

	return text
}
