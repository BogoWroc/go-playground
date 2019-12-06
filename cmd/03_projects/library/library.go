package library

type Author struct {
	FirstName string
	LastName  string
}

type Bookcase struct {
	no          int
	bookshelves []Bookshelf
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
