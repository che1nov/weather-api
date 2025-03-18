# Weather API Project

## Overview

This project involves building a weather API in Go that fetches and returns weather data from a 3rd party API. By working on this project, you will learn how to interact with 3rd party APIs, implement caching, and manage environment variables.

## API Selection

You can use any weather API of your choice. As a suggestion, you can use Visual Crossing’s API, which is free and easy to use.

## Caching

For caching, we recommend using Redis. You can use the city code entered by the user as the key and store the API response in Redis with an expiration time. This way, the cache will automatically clean itself when the data becomes outdated (e.g., after 12 hours).

## Tips for Implementation

1. **Start Simple**: Create a simple API that returns a hardcoded weather response. This will help you understand how to structure your API and handle requests.
2. **Environment Variables**: Use environment variables to store the API key and the Redis connection string. This allows for easy updates without modifying the code.
3. **Error Handling**: Ensure proper error handling. If the 3rd party API is down or the city code is invalid, return the appropriate error message.
4. **HTTP Requests**: Use a package or module to make HTTP requests. For Go, you can use the `net/http` package or a third-party package like `go-resty/resty`.
5. **Rate Limiting**: Implement rate limiting to prevent abuse of your API. You can use a package like `golang.org/x/time/rate`.

## Getting Started

### Prerequisites

- Go installed
- Redis installed and running
- Visual Crossing API key

### Installation

1. Clone the repository:
   ```sh
   git clone <repository_url>
   cd weather-api
   ```

2. Set up environment variables:
   Create a `.env` file and add your API key and Redis connection string:
   ```env
   API_KEY=your_visual_crossing_api_key
   REDIS_URL=your_redis_connection_string
   ```

### Running the API

1. Build the application:
   ```sh
   go build -o weather-api
   ```

2. Run the application:
   ```sh
   ./weather-api
   ```

### API Endpoints

- `GET /weather?city={city_code}`: Fetches weather data for the specified city code.

### Example Request

```
GET /weather?city=NewYork
```

### Example Response

```json
{
  "city": "New York",
  "temperature": "15°C",
  "description": "Partly Cloudy",
  "humidity": "60%",
  "wind_speed": "10 km/h"
}
```

## License

This project is licensed under the MIT License.

https://roadmap.sh/projects/weather-api-wrapper-service