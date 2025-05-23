package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func (note Note) Display() {
	fmt.Printf("Your Note titled \"%v\" has the following content: \n\n%v\n\nCreated at: %v", note.title, note.content, note.createdAt)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	// cara ini kalo mau tetap priva title, content dan createdAt-nya (sulit dan ribet)
	json, err := json.Marshal(&struct {
		Title     string    `json:"Judul"`
		Content   string    `json:"isi"`
		CreatedAt time.Time `json:"waktu"`
	}{
		Title:     note.title,
		Content:   note.content,
		CreatedAt: note.createdAt,
	})

	// cara gampangnya, ada warning "struct type 'example.com/note/note.Note' doesn't have any exported fields, nor custom marshaling (SA9005)" solusinya ubah ke huruf besar:
	// type Note struct {
	// Title     string    `json:"Judul"`
	// Content   string    `json:"isi"`
	// CreatedAt time.Time `json:"pembuatan"`
	// }

	// json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return nil, errors.New("title and content are required")
	}

	return &Note{
		title,
		content,
		time.Now(),
	}, nil
}
