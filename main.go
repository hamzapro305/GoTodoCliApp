package main

import (
	"TodoCli/Command"
	"TodoCli/CustomStorage"
	"TodoCli/Todo"
)

func main() {
	todos := Todo.Todos{}
	storage := CustomStorage.NewStorage[Todo.Todos]("todos.json")
	storage.Load(&todos)
	cmdFlags := Command.NewCmdFlags()
	cmdFlags.Execute(&todos)
	storage.Save(todos)
}
