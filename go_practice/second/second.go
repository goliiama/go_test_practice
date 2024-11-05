package second

import "fmt"

type User struct {
	Id   int
	Name string
}

type Book struct {
	ID         int
	Title      string
	Author     string
	IsBorrowed bool
}

type LibraryServices interface {
	AddBook(book Book) error
	ReturnBook(bookID int) error
	BorrowBook(bookID int, user User) error
	ListofBooks() []Book
}

type Library struct {
	books map[int]*Book
}

// creating an empthy library (constroctor function for the Library struct), in *Library, Library is not the struct we defined before, its an instance Library that we created inside the memory and now in this function we are  returning a pointer(address) to that library
func NewLibrary() *Library {
	return &Library{books: make(map[int]*Book)}

}

func (l *Library) AddBook(book Book) error {
	if _, exists := l.books[book.ID]; exists {
		return fmt.Errorf("book id with %v id already exists", book.ID)
	}
	l.books[book.ID] = &book
	return nil
}

func (l *Library) BorrowBook(bookID int, user User) error {
	book, exists := l.books[bookID]
	if !exists {
		return fmt.Errorf("book with id %v does not exists", bookID)
	}

	if book.IsBorrowed {
		return fmt.Errorf("book %v is already borroed", book.Title)
	}
	book.IsBorrowed = true
	return nil
}

func (l *Library) ReturnBook(bookID int) error {
	book, exists := l.books[bookID]
	if !exists {
		return fmt.Errorf("book with %v ID not found", bookID)
	}

	if !book.IsBorrowed {
		return fmt.Errorf("book with %v is already retured", bookID)
	}

	book.IsBorrowed = false
	return nil
}

func (l *Library) ListofBooks() []Book {
	availableBook := []Book{}
	for _, book := range l.books {
		if !book.IsBorrowed {
			availableBook = append(availableBook, *book)
		}
	}
	return availableBook
}
