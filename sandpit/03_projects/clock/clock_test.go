package clock

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Clock tests", func() {

	It("should return clock set on proper hour and minute", func() {
		clock := NewClock(0, 0)
		Expect(clock.String()).To(Equal("12:00"))
	})

	It("should return clock set on 12:05", func() {
		clock := NewClock(0, 0)
		clock.AddMinutes(5)
		Expect(clock.String()).To(Equal("12:05"))
	})

	It("should return clock set on 12:05", func() {
		clock := NewClock(0, 0)
		clock.AddMinutes(49)
		Expect(clock.String()).To(Equal("12:49"))
	})

	It("should return clock set on 01:00", func() {
		clock := NewClock(0, 0)
		clock.AddMinutes(60)
		Expect(clock.String()).To(Equal("01:00"))
	})

	It("should return clock set on 01:01", func() {
		clock := NewClock(0, 0)
		clock.AddMinutes(61)
		Expect(clock.String()).To(Equal("01:01"))
	})

	It("should return clock set on 01:59", func() {
		clock := NewClock(0, 0)
		clock.AddMinutes(119)
		Expect(clock.String()).To(Equal("01:59"))
	})

	It("should return clock set on 02:00", func() {
		clock := NewClock(0, 0)
		clock.AddMinutes(120)
		Expect(clock.String()).To(Equal("02:00"))
	})

	It("should return clock set on 11:10", func() {
		clock := NewClock(0, 0)
		clock.AddMinutes(670)
		Expect(clock.String()).To(Equal("11:10"))
	})

	It("should return clock set on 12:00", func() {
		clock := NewClock(0, 0)
		clock.AddMinutes(720)
		Expect(clock.String()).To(Equal("12:00"))
	})

	It("should return clock set on 01:00", func() {
		clock := NewClock(0, 0)
		clock.AddMinutes(780)
		Expect(clock.String()).To(Equal("01:00"))
	})
})
