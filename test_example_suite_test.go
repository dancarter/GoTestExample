package calculator_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTestExample(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestExample Suite")
}
