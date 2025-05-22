# url-shortener

A simple and efficient URL shortener service written in Go.

## Features

- Generate short, shareable URLs from long links
- Redirect users from short URLs to original destinations
- Minimal, fast, and easy to deploy
- Written in pure Go for portability and performance

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) 1.18 or later

### Installation

Clone the repository:

```bash
git clone https://github.com/lutffmn/url-shortener.git
cd url-shortener
```

Build the application:

```bash
go build -o url-shortener
```

### Running the Application

You can run the application with:

```bash
./url-shortener
```

By default, the service will start on port 8080. You can customize configuration via environment variables or command-line flags (add details here if available).

## Usage

- Send a POST request to `/shorten` with a JSON payload containing the URL you want to shorten.
- Access a shortened URL by visiting `http://localhost:8080/<shortcode>`.

Example curl command:

```bash
curl -X POST -H "Content-Type: application/json" \
  -d '{"url":"https://www.example.com"}' \
  http://localhost:8080/shorten
```

## API

Base URL:

```
/api/v1/
```

| Method | Endpoint              | Description                        |
| ------ | --------------------- | ---------------------------------- |
| POST   | `/shorten`            | Create a shortened URL             |
| GET    | `/{alias}`            | Redirect to original URL           |
| GET    | `/urls/{alias}`       | Get URL details                    |
| GET    | `/urls/`              | Get user's shortened URL list      |
| PATCH  | `/urls/{alias}`       | Update a shortened URL expire time |
| DELETE | `/urls/{alias}`       | Delete a shortened URL             |
| GET    | `/urls/{alias}/stats` | GET URL stats details              |

## Project Structure

- `cmd/api/server.go` - Application entry point
- `internal/handlers/` - HTTP request handlers
- `pkg/storage/` - URL storage logic

## Upcoming Features

**Soon**

## Contributing

Contributions are welcome! Please open issues or submit pull requests for bug fixes or new features.

## License

This project is licensed under the MIT License.

## Acknowledgments

- Inspired by popular URL shortener services
- Built with Go and love 💙

---
