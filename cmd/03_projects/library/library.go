package library

import "errors"

type Author struct {
	FirstName string
	LastName  string
}

type Bookcase struct {
	no          int
	bookshelves []Bookshelf
}

func (b Bookcase) isEnoughCapacityOnShelves(books []Book) bool {

	capacity := 0
	for _, bookshelf := range b.bookshelves {
		capacity += bookshelf.capacity
	}

	if capacity >= len(books) {
		return true
	}

	return false
}

func (b Bookcase) store(books []Book) {

}

type Bookshelf struct {
	no       int
	capacity int
	books    []Book
}

type Book struct {
	title  string
	author Author
}

type Library struct {
	bookcases []Bookcase
}

func (l Library) RegisterBooks(books []Book) (bool, error) {
	bookcase, err := l.findBookcaseFor(books)
	if err != nil {
		return false, errors.New("there is no free bookcase")
	}
	bookcase.store(books)
	return true, nil
}

func (l Library) findBookcaseFor(books []Book) (Bookcase, error) {
	for _, bookcase := range l.bookcases {
		if bookcase.isEnoughCapacityOnShelves(books) {
			return bookcase, nil
		}
	}
	return Bookcase{}, errors.New("no free bookcase")
}

// NewLibrary will create library with bookcases
func NewLibrary(noOfBookcases, noOfBookshelves, bookshelfCapacity int) Library {
	bookcases := make([]Bookcase, 0)
	for index := 0; index < noOfBookcases; index++ {
		bookshelves := createBookshelves(noOfBookshelves, bookshelfCapacity)
		bookcases = append(bookcases, Bookcase{
			no:          index,
			bookshelves: bookshelves,
		})
	}
	return Library{bookcases: bookcases}
}

func createBookshelves(noOfBookshelves int, bookshelfCapacity int) []Bookshelf {
	bookshelves := make([]Bookshelf, 0)
	for index := 0; index < noOfBookshelves; index++ {
		bookshelves = append(bookshelves, Bookshelf{
			no:       index,
			capacity: bookshelfCapacity,
			books:    make([]Book, 0),
		})
	}
	return bookshelves
}
