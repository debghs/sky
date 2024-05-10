# sky

sky is command line tool for fetching and displaying the current weather and forecast for a specified city using the WeatherAPI.

## Dependencies
This program relies on the following external dependencies:

- WeatherAPI: Provides weather data via an API.

## Configuration
  You need to obtain an API key from WeatherAPI. 
  <br>Once you have the key, you can set it in the main function of main.go:

  ```response, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=YOUR_API_KEY&q=" + query + "&days=1&aqi=no&alerts=no")```

## Installation

1. Clone the repository:

   ```git clone https://github.com/debghs/sky.git```
   
2. Navigate to the project directory:
   
   ```cd sky```
   
3. Build the Go program:
   
   ```go build```
   
## Usage
Run the program from the command line: ```./sky [city]```

  Replace [city] with the name of the city you want to get the weather forecast for.
  <br>If no city is provided, it defaults to "Kolkata".

  For example: ```./sky Paris```

  ## Features
  If there is more than 40% chance of rain, it highlights the entry in red color.
