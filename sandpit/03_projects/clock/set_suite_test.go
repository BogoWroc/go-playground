package clock

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestSet(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Set Suite")
}
