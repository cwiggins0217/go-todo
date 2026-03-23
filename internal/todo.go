package internal

import (
	"fmt"
	"time"
)

type Todo struct {
	Name        string
	Description string
	Id          int
	timestamp   time.Time
}

func NewTodo(name string, description string, lineCount int) (Todo, error) {
	// caveman error handling because I baby
	if name == "" {
		return Todo{}, fmt.Errorf("invalid name")
	}
	if description == "" {
		return Todo{}, fmt.Errorf("invalid description")
	}
	if lineCount < 0 {
		return Todo{}, fmt.Errorf("invalid line count")
	}

	return Todo{Name: name, Description: description, Id: lineCount + 1, timestamp: time.Now()}, nil
}
