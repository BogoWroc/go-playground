package storage

import (
	"errors"
	"fmt"
	"sync"

	. "github.com/bogowroc/go-playground/sandpit/03_projects/store/api"
)

var mutex sync.Mutex

type inMemoryStorage struct {
	storage []Product
	id      int
}

func (s *inMemoryStorage) SaveOrUpdate(product Product) (Product, error) {
	if product.Id == 0 {
		p := Product{
			Id:   s.generateId(),
			Name: product.Name,
		}
		s.storage = append(s.storage, p)

		return p, nil
	} else {
		for i, p := range s.storage {
			if p.Id == product.Id {
				pRef := &s.storage[i]
				pRef.Name = product.Name
				return *pRef, nil
			}
		}
		return product, errors.New(fmt.Sprintf("Unable to find product with id = %v", product.Id))
	}

}

func (s *inMemoryStorage) generateId() int {
	mutex.Lock()
	s.id += 1
	id := s.id
	mutex.Unlock()

	return id
}

func (s inMemoryStorage) FetchAll() []Product {
	return s.storage
}

func New() Storage {
	return &inMemoryStorage{
		id:      0,
		storage: make([]Product, 0),
	}
}
