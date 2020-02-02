package core

import (
	. "github.com/bogowroc/go-playground/sandpit/03_projects/lift/api"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type DisplayMock struct {
	message string
}

func (d *DisplayMock) Display(message string) {
	d.message = message
}

var _ = Describe("lift tests", func() {

	var l Lift
	var displayMock DisplayMock

	BeforeEach(func() {
		// mocks
		displayMock = DisplayMock{}

		l = New(150, &displayMock)
		// also we can redefine value by interface conversion
		//l.(*lift).display = &displayMock
	})

	It("should do not report that max weight was exceeded when lift is empty", func() {

		Expect(displayMock.message).To(Equal(""))

	})

	It("should do not report that max weight was exceeded", func() {

		p1 := Person{Weight: 90}
		p2 := Person{Weight: 40}
		l.Enter(p1, p2)

		Expect(displayMock.message).To(Equal(""))

	})

	It("should report that max weight exceeded", func() {

		p1 := Person{Weight: 90}
		p2 := Person{Weight: 80}
		l.Enter(p1, p2)

		Expect(displayMock.message).To(Equal("Max weight exceeded"))
	})

})
