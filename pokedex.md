# Pokedex CLI

The **Pokedex CLI** is an interactive terminal-based application that allows you to explore Pokémon locations, catch Pokémon, view your collection, and inspect your caught Pokémon.

## 📦 Features

- Browse Pokémon location areas using paginated API calls.
- Explore a location to list available Pokémon there.
- Catch Pokémon and store them in your personal pokedex.
- Inspect detailed stats for each caught Pokémon.
- View all caught Pokémon.
- Interactive CLI with helpful command guide.
- In-memory caching system to save time on repeated requests.

## 🚀 Getting Started

### Prerequisites

- Go installed (`go 1.20+`)

### Running the App

```
git clone https://github.com/ethereumvd/gokit-cli.git
cd pokedex
go build . && ./pokedex
```
You’ll be greeted with:
```
hello
>>>
```

## 💡 Available Commands

| Command                   | Description                                |
|---------------------------|--------------------------------------------|
| `help`                    | Prints the help menu                       |
| `exit`                    | Turns off pokedex                         |
| `map`                     | Lists location areas (paginated forward)  |
| `mapb`                    | Lists location areas (paginated backward) |
| `explore <location_area>` | Lists Pokémon in a specific location       |
| `catch <pokemon_name>`    | Attempts to catch a specified Pokémon      |
| `inspect <pokemon_name>`  | Shows data about caught Pokémon            |
| `pokedex`                 | Displays all caught Pokémon                |

## 📌 Example Usage
```
>>> map
>>> explore viridian-forest
>>> catch pikachu
>>> pokedex
>>> inspect pikachu
```
## 📁 Directory Structure
```
pokedex
├── catch.go
├── commands.go
├── exit.go
├── explore.go
├── go.mod
├── help.go
├── inspect.go
├── internal
│   ├── pokeapi
│   │   ├── location_area.go
│   │   ├── location_area_req.go
│   │   ├── pokeapi.go
│   │   ├── pokemon.go
│   │   └── pokemon_req.go
│   └── pokecache
│       └── pokecache.go
├── main.go
├── map_mapb.go
└── pokedex.go
```

## 🧠 Notes

- Some commands may require internet access for fetching live data from the PokéAPI.
- Pokémon catching logic may involve probabilities.
