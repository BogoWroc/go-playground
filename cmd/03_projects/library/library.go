package library

import (
	"errors"
)

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
nextBook:
	for _, book := range books {
		for index := range b.bookshelves {
			if b.bookshelves[index].capacity > 0 {
				b.bookshelves[index].Add(book)
				continue nextBook
			}
		}
	}
}

type Bookshelf struct {
	no       int
	capacity int
	books    []Book
}

func (b *Bookshelf) Add(book Book) {
	b.books = append(b.books, book)
	b.capacity -= 1
}

type Book struct {
	title  string
	author Author
}

type Library struct {
	bookcases []Bookcase
}

type BookPosition struct {
	bookcaseId  int
	bookshelfId int
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

func (l Library) FindBook(book Book) BookPosition {
	for _, bookcase := range l.bookcases {
		for _, bookshelf := range bookcase.bookshelves {
			for _, b := range bookshelf.books {
				if b == book {
					return BookPosition{bookcase.no, bookshelf.no}
				}
			}
		}
	}

	return BookPosition{-1, -1}
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
