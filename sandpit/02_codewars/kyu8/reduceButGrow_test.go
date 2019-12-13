package kyu8

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Example", func() {
	It("Fixed tests", func() {
		Expect(Grow([]int{1, 2, 3})).To(Equal(6))
		Expect(Grow([]int{4, 1, 1, 1, 4})).To(Equal(16))
		Expect(Grow([]int{2, 2, 2, 2, 2, 2})).To(Equal(64))
	})
})
