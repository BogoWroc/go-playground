package library

import (
	"testing"
)

func TestBuildLibrary(t *testing.T) {

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
