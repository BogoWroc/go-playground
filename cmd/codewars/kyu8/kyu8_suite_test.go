package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestKyu8(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Kyu8 Suite")
}
