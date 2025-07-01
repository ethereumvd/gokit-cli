# Weather CLI

The **Weather CLI** is a terminal-based weather app written in Go. It fetches and displays the current weather and hourly forecast for any specified city using the [WeatherAPI](https://www.weatherapi.com/).

## 📦 Features

- View current temperature and conditions for a city
- Display hourly forecast for today
- Color-coded rain alerts (blue if high rain probability)
- Reads API key securely from `.env` file

## 🚀 Getting Started

### Prerequisites

- Go installed (`go 1.20+`)
- A free API key from [weatherapi.com](https://www.weatherapi.com/)
- `.env` file with your API key:

```env
API_KEY=your_weather_api_key_here
```

### Running the App
```bash
git clone https://github.com/ethereumvd/gokit-cli.git
cd weather
echo 'API_KEY=your_key_here' > .env
go build . && ./weather [city]
```

If you don't provide a city name, it defaults to Bhubaneshwar.

## 💡 Available Arguments

| Argument     | Description                                                       |
|--------------|-------------------------------------------------------------------|
| `[city]`     | *(Optional)* Name of the city to get weather for                 |
|              | *(e.g. `./weather London` or just `./weather` for default city)* |

## 📌 Example Usage
```bash
./weather
./weather London
./weather Tokyo
```

## 📁 Directory Structure
```
weather
├── .env
├── go.mod
├── go.sum
├── main.go
├── weather
└── weather.go
```
## 🧠 Notes

- Requires an internet connection to access live weather data
- Rain chance above 50% is highlighted using color (via [fatih/color](github.com/fatih/color))
- Make sure not to commit your .env file or API key to public repositories
