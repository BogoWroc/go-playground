package kyu8

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Basic tests", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(Goals(0, 0, 0)).To(Equal(0))
		Expect(Goals(43, 10, 5)).To(Equal(58))
	})

})
