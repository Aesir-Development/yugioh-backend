# Yu-Gi-Oh Card API

This is a simple API built with Go and the Gin framework. It connects to a database to fetch information about Yu-Gi-Oh cards.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go (version 1.16 or later is recommended)
- Gin framework
- A MySQL database

### Installing

1. Clone the repository to your local machine.
2. Navigate to the project directory.
3. Run `go get` to download the necessary Go packages.
4. Update the `dbConnection` package with your database connection details.

## Running the application

To start the server, run:

```bash
go run main.go

The server will start on port 8080.
```

## Endpoints
GET /cards: Fetches a list of all cards.
GET /test: A test endpoint that fetches all table names from the database. This is a temporary endpoint for testing purposes and will be removed in the future.
## Contributing
Please read CONTRIBUTING.md for details on our code of conduct, and the process for submitting pull requests to us.

## License
This project is licensed under the MIT License - see the LICENSE.md file for details
