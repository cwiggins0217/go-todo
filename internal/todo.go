package internal

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Id          int
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) Add(title string) {
	todo := Todo{
		Title:     title,
		Id:        0, //TODO fix this to include the index of todos array
		Completed: false,
		CreatedAt: time.Now(),
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("Invalid index")
		fmt.Println(err)
		return err
	} // this could probably be implemented in main or something

	return nil
}

func (todos *Todos) Delete(id int) error {
	t := *todos

	if err := t.validateIndex(id); err != nil {
		return err
	}

	*todos = append(t[:id], t[id+1:]...)

	return nil

}

func (todos *Todos) Toggle(index int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	isCompleted := t[index].Completed

	if !isCompleted {
		completionTime := time.Now()
		t[index].CompletedAt = &completionTime
	}

	t[index].Completed = !isCompleted

	return nil
}

func (todos *Todos) Edit(index int, title string) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	t[index].Title = title

	return nil
}

func (todos *Todos) Print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("ID", "Title", "Created", "Completed", "Completed At")
	for index, t := range *todos {
		completed := "❌"
		completedAt := ""

		if t.Completed {
			completed = "✅"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}

		table.AddRow(
			strconv.Itoa(index), t.Title,
			t.CreatedAt.Format(time.RFC1123), completed, completedAt)
	}

	table.Render()
}
