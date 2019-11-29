package kyu8

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Basic tests", func() {
	It("26 dogs", func() {
		Expect(HowManyDalmatians(26)).To(Equal("More than a handful!"))
	})
	It("8 dogs", func() {
		Expect(HowManyDalmatians(8)).To(Equal("Hardly any"))
	})
	It("14 dogs", func() {
		Expect(HowManyDalmatians(14)).To(Equal("More than a handful!"))
	})
	It("80 dogs", func() {
		Expect(HowManyDalmatians(80)).To(Equal("Woah that's a lot of dogs!"))
	})
	It("100 dogs", func() {
		Expect(HowManyDalmatians(100)).To(Equal("Woah that's a lot of dogs!"))
	})
	It("50 dogs", func() {
		Expect(HowManyDalmatians(50)).To(Equal("More than a handful!"))
	})
	It("10 dogs", func() {
		Expect(HowManyDalmatians(10)).To(Equal("Hardly any"))
	})
	It("101 dogs", func() {
		Expect(HowManyDalmatians(101)).To(Equal("101 DALMATIONS!!!"))
	})

})
