# Todo CLI

The **Todo CLI** is a fast and minimal command-line task manager written in Go. It allows you to add, edit, delete, toggle, and list todos using simple flags — no complex UI, just productivity.

## 📦 Features

- Add new todos directly from the command line
- List all existing todos
- Edit any todo by index
- Mark todos as completed/incomplete (toggle)
- Delete todos by index
- Simple flag-based interface
- In-memory with optional persistence

## 🚀 Getting Started

### Prerequisites

- Go installed (`go 1.20+`)

### Running the App
```
git clone https://github.com/ethereumvd/gokit-cli.git
cd todo
go build . && ./todo [flags]
```
## 💡 Available Flags

| Flag       | Description                                                  |
|------------|--------------------------------------------------------------|
| `-add`     | Add a new todo (e.g. `-add "Buy groceries"`)                |
| `-edit`    | Edit a todo by index with new title (e.g. `-edit 0:Read book`) |
| `-del`     | Delete a todo by index (e.g. `-del 2`)                       |
| `-toggle`  | Toggle completion status of a todo by index (e.g. `-toggle 1`) |
| `-list`    | List all todos                                               |

## 📌 Example Usage
```
./todo -add "Read a book"
./todo -list
./todo -toggle 0
./todo -edit 0:"Read Go documentation"
./todo -del 0
```
## 📁 Directory Structure
```
todo
├── commands.go
├── go.mod
├── go.sum
├── main.go
├── save.go
└── todo.go
```

## 🧠 Notes

- Indexing starts from 0 (e.g. -edit 0:...)
