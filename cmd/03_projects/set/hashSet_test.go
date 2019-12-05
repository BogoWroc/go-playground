package set

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hash tests", func() {

	var set Set

	BeforeEach(func() {
		set = NewHashSet()
	})

	It("should return false when element does not exist in hashSet", func() {
		set.Add(2)

		Expect(set.Contains(1)).To(Equal(false))
		Expect(set.Contains(1, 2, 3)).To(Equal(false))
	})

	It("should return true when element exist in hashSet", func() {
		set.Add(1, 2, 3)

		Expect(set.Contains(1)).To(Equal(true))
		Expect(set.Contains(1, 2, 3)).To(Equal(true))
	})

	It("should be possible to determine hashSet length", func() {
		set.Add(1, 2, 3)

		Expect(set.Length()).To(Equal(3))
	})

	It("should be possible to clean hashSet", func() {
		set.Add(1, 2, 3)

		set.Clean()
		Expect(set.Length()).To(Equal(0))
	})

	It("should be possible iterate over hashSet elements", func() {
		set.Add(1, 2, 3)

		i := 0
		for k := range set.Iterator() {
			_ = k
			i++
		}
		Expect(i).To(Equal(set.Length()))
	})

	It("should be possible to convert hashSet to slice", func() {
		set.Add(1, 2, 3)

		Expect(len(set.ToSlice())).To(Equal(3))
	})

})
