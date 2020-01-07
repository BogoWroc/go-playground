package api

type Product struct {
	Id   int
	Name string
}

type Storage interface {
	SaveOrUpdate(product Product) (Product, error)
	FetchAll() []Product
}
