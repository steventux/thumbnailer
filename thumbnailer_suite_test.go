package thumbnailer_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestThumbnailer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Thumbnailer Suite")
}
