package thumbnailer_test

import (
	. "thumbnailer"

	"bytes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
)

var _ = Describe("Thumbnailer", func() {

	var res *httptest.ResponseRecorder

	var testFileContents = func() (os.FileInfo, []byte, error) {
		file, err := os.Open("./thumbnailer.go")
		if err != nil {
			return nil, nil, err
		}
		fileContents, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, nil, err
		}
		fi, err := file.Stat()
		if err != nil {
			return nil, nil, err
		}
		file.Close()
		return fi, fileContents, nil
	}

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
		It("should complain if a size is not supplied", func() {
			req, _ := http.NewRequest("POST", "http://localhost/thumbnail", nil)
			ThumbnailHandler(res, req)

			Expect(res.Body.String()).To(Equal("Please specify a thumbnail size eg. 100x100"))
		})
		Describe("Given sane POST data", func() {
			BeforeEach(func() {
				body := new(bytes.Buffer)
				writer := multipart.NewWriter(body)
				writer.WriteField("size", "100x100")
				fi, fileContents, _ := testFileContents()
				part, _ := writer.CreateFormFile("file", fi.Name())
				part.Write(fileContents)
				writer.Close()

				req, _ := http.NewRequest("POST", "http://localhost/thumbnail", body)
				req.Header.Add("Content-Type", writer.FormDataContentType())

				ThumbnailHandler(res, req)
			})

			It("should be happy if size is given", func() {
				Expect(res.Body.String()).To(MatchRegexp("Sizing to 100x100"))
			})
			It("should be happy if a file is supplied", func() {
				Expect(res.Body.String()).To(MatchRegexp("File uploaded successfully"))
			})
		})
	})
})
