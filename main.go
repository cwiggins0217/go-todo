package main

import "github.com/cwiggins0217/go-todo/internal"

func main() {
	todos := internal.Todos{}
	todos.Add("buy milk", "buy me some milk")
	todos.Add("buy bread", "buy me some bread")
	todos.Toggle(0)
	todos.Print()
}
