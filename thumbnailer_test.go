package thumbnailer_test

import (
	. "thumbnailer"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"strings"
)

var _ = Describe("Thumbnailer", func() {

	var res *httptest.ResponseRecorder

	BeforeEach(func() {
		res = httptest.NewRecorder()
	})

	Describe("RootHandler", func() {
		It("should say Welcome", func() {
			req, _ := http.NewRequest("GET", "http://localhost/", nil)
			RootHandler(res, req)

			Expect(res.Body.String()).To(Equal("Welcome to Thumbnailer"))
		})
	})

	Describe("ThumbnailHandler", func() {
		It("should complain about GET requests", func() {
			req, _ := http.NewRequest("GET", "http://localhost/thumbnail/100x100", nil)
			ThumbnailHandler(res, req)

			Expect(res.Body.String()).To(Equal("I only handle POSTs"))
		})
		It("should complain if required params are not passed", func() {
			req, _ := http.NewRequest("POST", "http://localhost/thumbnail", nil)
			ThumbnailHandler(res, req)

			Expect(res.Body.String()).To(Equal("You've forgotten something"))
		})
		It("should be happy if required params are passed", func() {
			req, _ := http.NewRequest("POST", "http://localhost/thumbnail", strings.NewReader("size=100x100"))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

			ThumbnailHandler(res, req)

			Expect(res.Body.String()).To(Equal("Sizing to 100x100"))
		})
	})
})
