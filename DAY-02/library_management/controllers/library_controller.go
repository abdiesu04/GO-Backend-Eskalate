package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"library_management/models"
	"library_management/services"
)

type LibraryController struct {
	libraryService services.LibraryManager
}

func NewLibraryController() *LibraryController {
	return &LibraryController{
		libraryService: services.NewLibraryService(),
	}
}

func (lc *LibraryController) AddBook() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter book ID: ")
	scanner.Scan()
	bookID, _ := strconv.Atoi(scanner.Text())
	fmt.Print("Enter book title: ")
	scanner.Scan()
	title := scanner.Text()
	fmt.Print("Enter book author: ")
	scanner.Scan()
	author := scanner.Text()

	book := models.Book{ID: bookID, Title: title, Author: author, Status: "Available"}
	lc.libraryService.AddBook(book)
	fmt.Println("\n✅ Book added successfully!\n")
}

func (lc *LibraryController) RemoveBook() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter book ID to remove: ")
	scanner.Scan()
	bookID, _ := strconv.Atoi(scanner.Text())

	err := lc.libraryService.RemoveBook(bookID)
	if err != nil {
		fmt.Printf("\n❌ Error: %s\n\n", err)
	} else {
		fmt.Println("\n✅ Book removed successfully!\n")
	}
}

func (lc *LibraryController) BorrowBook() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter book ID to borrow: ")
	scanner.Scan()
	bookID, _ := strconv.Atoi(scanner.Text())
	fmt.Print("Enter member ID: ")
	scanner.Scan()
	memberID, _ := strconv.Atoi(scanner.Text())

	err := lc.libraryService.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Printf("\n❌ Error: %s\n\n", err)
	} else {
		fmt.Println("\n✅ Book borrowed successfully!\n")
	}
}

func (lc *LibraryController) ReturnBook() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter book ID to return: ")
	scanner.Scan()
	bookID, _ := strconv.Atoi(scanner.Text())
	fmt.Print("Enter member ID: ")
	scanner.Scan()
	memberID, _ := strconv.Atoi(scanner.Text())

	err := lc.libraryService.ReturnBook(bookID, memberID)
	if err != nil {
        fmt.Printf("\n❌ Error: %s\n\n", err)
        } else {
            fmt.Println("\n✅ Book returned successfully!\n")
        }
    }
    func (lc *LibraryController) ListAvailableBooks() {
        books := lc.libraryService.ListAvailableBooks()
        if len(books) == 0 {
            fmt.Println("\nNo available books at the moment.\n")
            return
        }
    
        fmt.Println("\n📚 Available Books:")
        for _, book := range books {
            fmt.Printf("- ID: %d, Title: \"%s\", Author: %s\n", book.ID, book.Title, book.Author)
        }
        fmt.Println()
    }


func (lc *LibraryController) ListBorrowedBooks() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter member ID: ")
	scanner.Scan()
	memberID, _ := strconv.Atoi(scanner.Text())

	books := lc.libraryService.ListBorrowedBooks(memberID)
	if len(books) == 0 {
		fmt.Println("\nThis member has not borrowed any books.\n")
		return
	}

	fmt.Println("\n📚 Borrowed Books:")
	for _, book := range books {
		fmt.Printf("- ID: %d, Title: \"%s\", Author: %s\n", book.ID, book.Title, book.Author)
	}
	fmt.Println()
}
