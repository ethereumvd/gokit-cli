package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title string
	Completed bool
	CreatedAt time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) AddTodo(title string) {

	NewTodo := Todo{
		Title: title,
		Completed: false,
		CreatedAt: time.Now(),
		CompletedAt: nil,
	}

	*todos = append(*todos, NewTodo)
}

func (todos *Todos) IsValidIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		return fmt.Errorf("Index is invalid")
	}

	return nil
}

func (todos * Todos) DeleteTodo(index int) error {

	err := todos.IsValidIndex(index)
	if err != nil { return err }

	DummyTodos := *todos
	*todos = append(DummyTodos[:index], DummyTodos[index+1:]...)

	return nil
}


func (todos *Todos) ToggleTodo(index int) error {

	t := *todos

	err := todos.IsValidIndex(index)
	if err != nil { return err }

	isCompleted := t[index].Completed

	if !isCompleted {
		CompletionTime := time.Now()
		t[index].CompletedAt = &CompletionTime
	}

	t[index].Completed = !isCompleted

	return nil
}


func (todos *Todos) EditTodo(index int, title string) error {

	t := *todos

	err := todos.IsValidIndex(index)
	if err != nil { return err }

	t[index].Title = title

	return nil
}

func (todos *Todos) PrintTodos() {

	table := table.New(os.Stdout)

	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")

	for index , todo := range *todos {

		completed := "❌"
		completedAt := ""
		createdAt := todo.CreatedAt.Format(time.RFC1123)

		if todo.Completed {
			completed = "✅"
			if todo.CompletedAt != nil {
				completedAt = todo.CompletedAt.Format(time.RFC1123)
			}
		}

		table.AddRow(strconv.Itoa(index+1), todo.Title, completed, createdAt, completedAt)

	}
	table.Render()

}