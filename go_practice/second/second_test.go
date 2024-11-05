package second

import (
	"testing"
)

func TestAddBook(t *testing.T) {
	lib := NewLibrary()
	book := Book{ID: 1, Title: "Book1", Author: "Goli", IsBorrowed: false}
	err := lib.AddBook(book)
	if err != nil {
		t.Fatalf("expected no error, we got %v", err)
	}
}

func TestBorrowBook(t *testing.T) {
	lib := NewLibrary()
	book := Book{ID: 1, Title: "book1", Author: "goli", IsBorrowed: false}
	lib.AddBook(book)

	err := lib.BorrowBook(1, User{Id: 1, Name: "sara"})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	err = lib.BorrowBook(1, User{1, "naz"})
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

}

func TestReturnBook(t *testing.T) {
	lib := NewLibrary()
	book := Book{ID: 1, Title: "book1", Author: "goli", IsBorrowed: false}
	lib.AddBook(book)
	lib.BorrowBook(1, User{1, "sara"})
	err := lib.ReturnBook(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestListAvailableBooks(t *testing.T) {
	lib := NewLibrary()
	book := Book{ID: 1, Title: "book1", Author: "goli", IsBorrowed: false}
	lib.AddBook(book)
	bookList := lib.ListofBooks()
	if len(bookList) != 1 {
		t.Fatalf("expected 1 availabality, got %v", len(bookList))
	}
}
