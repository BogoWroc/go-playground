package kyu8

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Example", func() {
	It("tests basic strings", func() {
		Expect(IsPalindrome("a")).To(Equal(true))
		Expect(IsPalindrome("aba")).To(Equal(true))
		Expect(IsPalindrome("Abba")).To(Equal(true))
		Expect(IsPalindrome("hello")).To(Equal(false))
	})
})
