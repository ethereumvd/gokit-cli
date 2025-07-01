package main

type Config struct {
	todos Todos
	storage Storage[Todos]
	cliflags *CmdFlags
}

func main() {

	cfg := Config{
		todos: Todos{},
	}

	cfg.storage = *NewStorage[Todos]("todos.json")
	cfg.storage.LoadData(&cfg.todos)
	cfg.cliflags = NewCLIFlags()
	cfg.cliflags.Execute(&cfg.todos)
	cfg.storage.SaveStorage(cfg.todos)
}
