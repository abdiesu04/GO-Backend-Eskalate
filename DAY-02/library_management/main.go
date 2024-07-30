package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"library_management/controllers"
)

func main() {
	libraryController := controllers.NewLibraryController()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Library Management System")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books by Member")
		fmt.Println("7. Exit")
		fmt.Print("Select an option: ")

		scanner.Scan()
		option, _ := strconv.Atoi(scanner.Text())

		switch option {
		case 1:
			libraryController.AddBook()
		case 2:
			libraryController.RemoveBook()
		case 3:
			libraryController.BorrowBook()
		case 4:
			libraryController.ReturnBook()
		case 5:
			libraryController.ListAvailableBooks()
		case 6:
			libraryController.ListBorrowedBooks()
		case 7:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option, please try again.")
		}

		fmt.Println("Do you want to continue? (y/n): ")
		scanner.Scan()
		choice := scanner.Text()
		if choice != "y" && choice != "Y" {
			fmt.Println("Goodbye!")
			break
		}
	}
}
