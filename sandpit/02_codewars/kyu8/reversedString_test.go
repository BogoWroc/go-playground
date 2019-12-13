package kyu8

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Basic tests", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(ReversedString("world")).To(Equal("dlrow"))
	})

})
