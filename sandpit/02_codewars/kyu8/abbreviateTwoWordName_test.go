package kyu8

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Example", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(AbbrevName("Sam Harris")).To(Equal("S.H"))
		Expect(AbbrevName("Patrick Feenan")).To(Equal("P.F"))
		Expect(AbbrevName("Evan Cole")).To(Equal("E.C"))
		Expect(AbbrevName("P Favuzzi")).To(Equal("P.F"))
		Expect(AbbrevName("David Mendieta")).To(Equal("D.M"))
	})
})
