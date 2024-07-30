package services

import (
	"errors"
	"sync"

	"library_management/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int) error
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

type LibraryService struct {
	books   map[int]models.Book
	members map[int]models.Member
	mu      sync.Mutex
}

func NewLibraryService() LibraryManager {
	return &LibraryService{
		books:   make(map[int]models.Book),
		members: make(map[int]models.Member),
	}
}

func (ls *LibraryService) AddBook(book models.Book) {
	ls.mu.Lock()
	defer ls.mu.Unlock()
	ls.books[book.ID] = book
}

func (ls *LibraryService) RemoveBook(bookID int) error {
	ls.mu.Lock()
	defer ls.mu.Unlock()
	if _, exists := ls.books[bookID]; exists {
		delete(ls.books, bookID)
		return nil
	}
	return errors.New("book not found")
}

func (ls *LibraryService) BorrowBook(bookID int, memberID int) error {
	ls.mu.Lock()
	defer ls.mu.Unlock()
	book, bookExists := ls.books[bookID]
	member, memberExists := ls.members[memberID]

	if !bookExists {
		return errors.New("book not found")
	}

	if book.Status == "Borrowed" {
		return errors.New("book already borrowed")
	}

	if !memberExists {
		member = models.Member{ID: memberID}
	}

	book.Status = "Borrowed"
	ls.books[bookID] = book
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	ls.members[memberID] = member
	return nil
}

func (ls *LibraryService) ReturnBook(bookID int, memberID int) error {
	ls.mu.Lock()
	defer ls.mu.Unlock()
	book, bookExists := ls.books[bookID]
	member, memberExists := ls.members[memberID]

	if !bookExists {
		return errors.New("book not found")
	}

	if !memberExists {
		return errors.New("member not found")
	}

	bookFound := false
	for i, borrowedBook := range member.BorrowedBooks {
		if borrowedBook.ID == bookID {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			bookFound = true
			break
		}
	}

	if !bookFound {
		return errors.New("book not borrowed by member")
	}

	book.Status = "Available"
	ls.books[bookID] = book
	ls.members[memberID] = member
	return nil
}

func (ls *LibraryService) ListAvailableBooks() []models.Book {
	ls.mu.Lock()
	defer ls.mu.Unlock()
	var availableBooks []models.Book
	for _, book := range ls.books {
		if book.Status == "Available" {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

func (ls *LibraryService) ListBorrowedBooks(memberID int) []models.Book {
	ls.mu.Lock()
	defer ls.mu.Unlock()
	member, exists := ls.members[memberID]
	if !exists {
		return nil
	}
	return member.BorrowedBooks
}
