# Learning Projects: Go Bookstore, Go Movies CRUD, Go Server

This repository contains three learning projects I completed while following a crash course on Go programming by freeCodeCamp.org. The three projects are:

- `go-server`: A HTTP server handling simple GET and POST requests.
- `go-movies-crud`: A CRUD (Create, Read, Update, Delete) web application that allows users to manage a list of movies.
- `go-bookstore`: A web application with a mysql db that allows users to browse, edit, add, delete and purchase books.


## Getting Started

To get started with these projects, first clone the repository to your local machine:

```
git clone https://github.com/your-username/go-learning-projects.git
```

Each project is contained in its own directory (`go-bookstore`, `go-movies-crud`, and `go-server`). To run a project, navigate to its directory and run the `main.go` file using the `go run` command.

For example, to run the `go-bookstore` project:

```
cd go-bookstore
go run main.go
```

This will start the `go-bookstore` web application on `http://localhost:8080`.

## Dependencies

These projects use the following third-party dependencies:

- `gorilla/mux`: A powerful HTTP router and URL matcher for building web applications in Go.
- `gorm`: A powerful ORM (Object-Relational Mapping) library for Go that provides a high-level interface for interacting with databases.
- `go-sqlite3`: A driver for the SQLite database that allows Go programs to interact with SQLite databases.

You can install these dependencies using the `go get` command:

```
go get -u github.com/gorilla/mux
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
```
