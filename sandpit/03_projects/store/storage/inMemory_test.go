package storage

import (
	. "github.com/bogowroc/go-playground/sandpit/03_projects/store/api"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("inMemoryStorage tests", func() {

	var inMemoryStorage Storage

	BeforeEach(func() {
		inMemoryStorage = New()
	})

	It("should store product in memory", func() {
		p := Product{
			Id:   0,
			Name: "Test",
		}

		p, err := inMemoryStorage.SaveOrUpdate(p)

		Expect(err).To(BeNil())
		Expect(p.Id).To(Equal(1))
	})

	It("should report an error when product cannot be found during instance update", func() {
		p := Product{
			Id:   1,
			Name: "Test",
		}
		product, err := inMemoryStorage.SaveOrUpdate(p)

		Expect(err.Error()).To(Equal("Unable to find product with id = 1"))
		Expect(product).To(Equal(p))
	})

	It("should update product in memory", func() {
		p1 := Product{
			Id:   0,
			Name: "Test 1",
		}

		p2 := Product{
			Id:   0,
			Name: "Test 2",
		}

		p3 := Product{
			Id:   0,
			Name: "Test 3",
		}

		inMemoryStorage.SaveOrUpdate(p1)
		inMemoryStorage.SaveOrUpdate(p2)
		product, _ := inMemoryStorage.SaveOrUpdate(p3)

		product.Name = "New value"

		p, err := inMemoryStorage.SaveOrUpdate(product)
		Expect(err).To(BeNil())
		Expect(p.Id).To(Equal(3))
		Expect(p.Name).To(Equal("New value"))
		Expect(inMemoryStorage.FetchAll()).To(ContainElement(Product{
			Id:   3,
			Name: "New value",
		}))
	})

})
