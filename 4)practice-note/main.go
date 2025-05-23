package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("saving the note failed")
		return
	}

	fmt.Println("saving the note succeeded!")

}

func getNoteData() (string, string) {
	title := getUserData("Note title: ")
	content := getUserData("Note content: ")

	return title, content
}

func getUserData(promptText string) string {
	fmt.Print(promptText)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text := strings.TrimSpace(input)

	return text
}
