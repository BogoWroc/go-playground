package kyu8

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestKyu8(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Kyu8 Suite")
}
