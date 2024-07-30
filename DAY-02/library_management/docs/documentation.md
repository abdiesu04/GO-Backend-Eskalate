# Library Management System

## Overview
This project is a simple console-based Library Management System implemented in Go. It demonstrates the use of structs, interfaces, and other Go functionalities such as methods, slices, and maps. The system allows you to add and remove books, borrow and return books, and list available and borrowed books.

## Features
- Add a new book to the library.
- Remove an existing book from the library.
- Borrow a book if it is available.
- Return a borrowed book.
- List all available books in the library.
- List all books borrowed by a specific member.

## Folder Structure

```
library_management/
├── main.go
├── controllers/
│   └── library_controller.go
├── models/
│   └── book.go
│   └── member.go
├── services/
│   └── library_service.go
├── docs/
│   └── documentation.md
└── go.mod
```

- `main.go`: Entry point of the application.
- `controllers/library_controller.go`: Handles console input and invokes the appropriate service methods.
- `models/book.go`: Defines the Book struct.
- `models/member.go`: Defines the Member struct.
- `services/library_service.go`: Contains business logic and data manipulation functions.
- `docs/documentation.md`: Contains system documentation and other related information.
- `go.mod`: Defines the module and its dependencies.

## Usage

1. **Add a new book**
    - Adds a new book to the library.
2. **Remove an existing book**
    - Removes a book from the library by its ID.
3. **Borrow a book**
    - Allows a member to borrow a book if it is available.
4. **Return a book**
    - Allows a member to return a borrowed book.
5. **List all available books**
    - Lists all available books in the library.
6. **List all borrowed books by a member**
    - Lists all books borrowed by a specific member.
