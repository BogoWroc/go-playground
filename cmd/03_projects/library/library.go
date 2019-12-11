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

func (b Bookcase) store(books ...Book) {
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

func (b Bookcase) freeSlots() (capacity int) {
	for _, bookshelf := range b.bookshelves {
		capacity += bookshelf.capacity
	}

	return
}

type Bookshelf struct {
	no       int
	capacity int
	books    []Book
}

func (b *Bookshelf) Add(book Book) {
	b.books = append(b.books, book)
	b.capacity--
}

type Book struct {
	title  string
	author Author
}

type Library struct {
	bookcases []Bookcase
}

type BookPosition struct {
	bookcaseID  int
	bookshelfID int
}

func (l Library) RegisterBooks(books []Book) (bool, error) {
	bookcases := l.findBookcasesFor(books)

	if len(bookcases) == 0 {
		return false, errors.New("there is no free bookcase")
	}
	for _, book := range books {
		for _, bookcase := range bookcases {
			if bookcase.freeSlots() > 0 {
				bookcase.store(book)
				break
			}
		}
	}
	return true, nil
}

func (l Library) findBookcasesFor(books []Book) []Bookcase {
	requiredCapacity := len(books)
	bookcases := make([]Bookcase, 0)
	for _, bookcase := range l.bookcases {
		freeSlots := bookcase.freeSlots()
		if freeSlots != 0 && requiredCapacity > 0 {
			bookcases = append(bookcases, bookcase)
		}

	}
	return bookcases
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
