package kyu8_test

import (
	"github.com/bogowroc/go-playground/pkg/codewars/kyu8"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Basic tests", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(kyu8.ReversedString("world")).To(Equal("dlrow"))
	})

})
