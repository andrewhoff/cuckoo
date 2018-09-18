package pure_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPure(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pure Suite")
}
