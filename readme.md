# GoKit CLI Tools 🛠️

A collection of fast and minimal command-line tools written in Go — built for productivity, learning, and fun. Each tool is self-contained and solves a specific real-world problem using clean Go code and idiomatic design.

## 🧩 Included Tools

| Tool       | Description                                                      |
|------------|------------------------------------------------------------------|
| 🧾 Todo CLI    | A minimal command-line todo/task manager                      |
| 🎮 Pokedex CLI | An interactive Pokédex explorer with caching and Pokémon catching |
| ☀️ Weather CLI | A weather app that shows current and hourly forecast for any city  |

---

## 🚀 Getting Started

### Prerequisites

- Go installed (`go 1.20+`)
- Internet connection (for Pokedex and Weather tools)
- [WeatherAPI](https://weatherapi.com/) API key for Weather CLI

### Clone the Repository

```bash
git clone https://github.com/ethereumvd/gokit-cli.git
cd gokit-cli
```

Each subproject is self-contained with its own go.mod, so you can build/run them independently.


## 📂 Directory Structure
```
gokit-cli/
├── pokedex/    # Explore, catch, and inspect Pokémon
├── todo/       # Manage todos from the command line
├── weather/    # View real-time weather and forecast info
├── readme.md   
├── weather.md  # Documentation for weather tool
├── todo.md     # Documentation for todo tool
└── pokedex.md  # Documentation for pokedex tool
```


## 📘 Project Highlights
### ✅ Clean Architecture

- Each tool lives in its own directory with main.go and modular code
- Organized packages with clear separation of concerns

### ⚙️ Real-World Integrations

- pokedex integrates with PokéAPI
- weather uses WeatherAPI with .env API key loading

### ✨ Terminal UX

- Interactive REPL-like experience in pokedex
- Flag-based control in todo and weather
- Colored CLI output for better readability


## 🔗 Related Readmes

- 🎮 [Pokedex CLI](https://github.com/ethereumvd/gokit-cli/blob/main/pokedex.md)

- 🧾 [Todo CLI](https://github.com/ethereumvd/gokit-cli/blob/main/todo.md)

- ☀️ [Weather CLI](https://github.com/ethereumvd/gokit-cli/blob/main/weather.md)

