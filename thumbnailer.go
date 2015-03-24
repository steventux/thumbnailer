package thumbnailer

import (
	//	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nfnt/resize"
	"image/jpeg"
	"path/filepath"
	//	"launchpad.net/goamz/aws"
	//	"launchpad.net/goamz/s3"
	//	"bufio"
	"io"
	"io/ioutil"
	//"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", RootHandler)
	r.HandleFunc("/thumbnail", RootHandler).Methods("GET")
	r.HandleFunc("/thumbnail", ThumbnailHandler).Methods("POST")
	http.Handle("/", r)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to Thumbnailer")
}

func ThumbnailHandler(w http.ResponseWriter, r *http.Request) {

	size := r.FormValue("size")
	if size == "" {
		fmt.Fprint(w, "Please specify a thumbnail size eg. 100x100")
		return
	}

	fmt.Fprint(w, "Sizing to "+size)

	// the FormFile function takes in the POST input id file
	file, header, err := r.FormFile("file")
	if err != nil {
		fmt.Fprint(w, err)
	}

	defer file.Close()

	tmpFile, err := ioutil.TempFile("", header.Filename)
	if err != nil {
		fmt.Fprint(w, err)
	}

	// write the content from POST to the file
	_, err = io.Copy(tmpFile, file)
	if err != nil {
		fmt.Fprintln(w, err)
	}

	tmpFile.Close()

	fmt.Fprintf(w, "File uploaded successfully : ")
	fmt.Fprintf(w, header.Filename)

	tmpFile, _ = os.Open(tmpFile.Name())
	thumbnailFile := thumbnail(tmpFile, size)

	fmt.Fprintf(w, "Thumbnail generated : ")
	fmt.Fprintf(w, thumbnailFile.Name())
}

func thumbnail(file *os.File, dimensions string) *os.File {

	width, height := parseDimensions(dimensions)

	img, err := jpeg.Decode(file)
	if err != nil {
		panic(err)
	}

	m := resize.Thumbnail(width, height, img, resize.Lanczos3)

	out, err := os.Create(thumbnailPath(file.Name(), width))
	if err != nil {
		panic(err)
	}

	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)

	return out
}

func parseDimensions(dimensions string) (uint, uint) {
	dimArr := strings.Split(dimensions, "x")
	width, _ := strconv.ParseUint(dimArr[0], 10, 64)
	height, _ := strconv.ParseUint(dimArr[1], 10, 64)
	return uint(width), uint(height)
}

func thumbnailPath(originalPath string, width uint) string {
	dir, file := filepath.Split(originalPath)
	ext := filepath.Ext(file)
	return filepath.Join(dir, file+"-thumb-"+strconv.FormatUint(uint64(width), 10)+ext)
}
