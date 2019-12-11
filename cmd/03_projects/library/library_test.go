package library

import (
	"strconv"
	"testing"
)

func TestCreateLibrary(t *testing.T) {

	tests := []struct {
		name                         string
		howManyBookcases             int
		howManyBookshelvesInBookcase int
		bookshelfCapacity            int
	}{
		{
			name:                         "shouldCreateLibrary",
			howManyBookcases:             50,
			howManyBookshelvesInBookcase: 3,
			bookshelfCapacity:            5,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			library := NewLibrary(
				test.howManyBookcases,
				test.howManyBookshelvesInBookcase,
				test.bookshelfCapacity)
			noOfBookcases := len(library.bookcases)
			if noOfBookcases != test.howManyBookcases {
				t.Errorf(
					"Invalid number of bookcases. Expected %d, but got %d",
					noOfBookcases,
					test.howManyBookcases)
			}

			noOfBookshelves := len(library.bookcases[0].bookshelves)
			if noOfBookshelves != test.howManyBookshelvesInBookcase {
				t.Errorf(
					"Invalid number of bookshelves. Expected %d, but got %d",
					noOfBookshelves,
					test.howManyBookshelvesInBookcase)
			}

			noOfBooks := len(library.bookcases[0].bookshelves[0].books)
			if noOfBooks != 0 {
				t.Errorf(
					"Invalid number of books at bookshelf. Expected %d, but got %d",
					noOfBooks,
					test.bookshelfCapacity)
			}
		})
	}
}

func TestRegisterBooksInLibrary(t *testing.T) {

	tests := []struct {
		name                string
		library             Library
		books               []Book
		acceptedByLibrarian bool
	}{
		{
			name:                "should not accept books when there are no bookcases in library",
			library:             NewLibrary(0, 0, 0),
			books:               giveBooks(2),
			acceptedByLibrarian: false,
		},
		{
			name:                "should not accept books when there is no place on bookshelves in any bookcase",
			library:             NewLibrary(1, 2, 0),
			books:               giveBooks(2),
			acceptedByLibrarian: false,
		},
		{
			name:                "should accept books when there is a capacity on bookshelves in any bookcase",
			library:             NewLibrary(1, 2, 2),
			books:               giveBooks(2),
			acceptedByLibrarian: true,
		},
		{
			name:                "should accept books when there is a capacity on bookshelves in any bookcase - two bookshelves will be used",
			library:             NewLibrary(1, 2, 2),
			books:               giveBooks(3),
			acceptedByLibrarian: true,
		},
		{
			name:                "should accept books when there is a capacity on bookshelves in any bookcase - two bookcases will be used",
			library:             NewLibrary(2, 2, 2),
			books:               giveBooks(6),
			acceptedByLibrarian: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if ok, _ := test.library.RegisterBooks(test.books); ok != test.acceptedByLibrarian {
				t.Errorf("Expected %v, got %v", test.acceptedByLibrarian, ok)
			}
		})
	}
}

func TestFindBook(t *testing.T) {

	library := NewLibrary(2, 2, 2)
	books := giveBooks(6)
	if ok, err := library.RegisterBooks(books); !ok {
		t.Error(err)
	}

	tests := []struct {
		name         string
		library      Library
		book         Book
		bookPosition BookPosition
	}{
		{
			name:         "should return number of bookcase and bookshelf - first bookcase",
			library:      library,
			book:         books[2],
			bookPosition: BookPosition{bookcaseID: 0, bookshelfID: 1},
		},
		{
			name:         "should return number of bookcase and bookshelf - second bookcase",
			library:      library,
			book:         books[5],
			bookPosition: BookPosition{bookcaseID: 1, bookshelfID: 0},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if bookPosition := test.library.FindBook(test.book); bookPosition != test.bookPosition {
				t.Errorf("Expected %v, got %v", test.bookPosition, bookPosition)
			}
		})
	}
}

func giveBooks(noOfBooks int) []Book {
	books := make([]Book, 0)
	for i := 0; i < noOfBooks; i++ {
		books = append(books, Book{
			title:  "Title: " + strconv.Itoa(i),
			author: Author{FirstName: "Bob", LastName: "Kent"},
		})
	}

	return books
}
