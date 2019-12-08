package kyu6

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func doTest(prod string, exp int) {
	var ans = duplicateCount(prod)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {

	It("should handle basic cases", func() {
		doTest("abcde", 0)
		doTest("abcdea", 1)
		doTest("abcdeaB11", 3)
		doTest("indivisibility", 1)
	})
})
