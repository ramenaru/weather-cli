## Usage
Here is an example weather report:
```bash
[ramenaru@arch-based weather-cli]$ go run main.go
? Enter the location (e.g., city): bandung

🌦️ Weather Forecast for Bandung, Indonesia
🌡️ Current Temperature: 32°C, Condition: Light drizzle

+-------+------------------+--------------------+----------------------+
| TIME  | TEMPERATURE (°C) | CHANCE OF RAIN (%) |      CONDITION       |
+-------+------------------+--------------------+----------------------+
| 16:00 | 31°C             |                 66 | Patchy rain possible |
| 17:00 | 29°C             |                 66 | Patchy rain possible |
| 18:00 | 27°C             |                 92 | Patchy rain possible |
| 19:00 | 26°C             |                 94 | Light rain shower    |
| 20:00 | 26°C             |                 94 | Light drizzle        |
| 21:00 | 27°C             |                 64 | Light rain shower    |
| 22:00 | 26°C             |                 89 | Patchy rain possible |
| 23:00 | 26°C             |                  0 | Partly cloudy        |
+-------+------------------+--------------------+----------------------+
```   

### Dependency
- Go 
- Git
- Bash

### Install dependencies
```bash
go mod download
go mod tidy
```    

### Running the script locally

Run the command `go run main.go`
