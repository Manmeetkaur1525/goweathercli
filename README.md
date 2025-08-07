# Weather CLI

A simple and clean command-line tool written in Go to fetch the current day's weather .

---

## Features

- Get today's weather from your terminal
- Lightweight and fast
- Easy to set up and use

---

## Installation

### 1. Clone the repository

```bash
git clone https://github.com/Manmeetkaur1525/goweathercli.git
cd goweathercli
```

### 2. Set your WeatherAPI key

```bash
export API_KEY=your_weatherapi_key
```

### 3. Build the CLI

```bash
go build -o weathercli
```

### 4. Move the binary to `/usr/local/bin` and rename it

```bash
sudo mv weathercli /usr/local/bin
```

---

## Usage

Now you can run it from anywhere:

```bash
today <city_name>
```

**Example:**

```bash
today Delhi
```

This will output the current day's weather for the given city.

---

## Requirements

- Go (version 1.16 or later)
- WeatherAPI key from [weatherapi.com](https://www.weatherapi.com)

---

## License

MIT
