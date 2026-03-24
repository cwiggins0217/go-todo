package main

import (
	"github.com/cwiggins0217/go-todo/cmd"
	"github.com/cwiggins0217/go-todo/internal"
)

func main() {
	todos := internal.Todos{}
	storage := internal.NewStorage[internal.Todos]("todos.json")
	cmdFlags := cmd.NewCmdFlags()
	cmdFlags.Execute(&todos)
	storage.Load(&todos)
	todos.Print()
	storage.Save(todos)
}
