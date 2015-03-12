package thumbnailer

import (
	//	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	//	"github.com/nfnt/resize"
	//	"image/jpeg"
	//	"launchpad.net/goamz/aws"
	//	"launchpad.net/goamz/s3"
	//	"log"
	"io"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", RootHandler)
	r.HandleFunc("/thumbnail", RootHandler).Methods("GET")
	r.HandleFunc("/thumbnail/{size}", ThumbnailHandler).Methods("POST")
	http.Handle("/", r)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)

	if err != nil {
		panic(err)
	}
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to Thumbnailer")
}

// TODO: Make this less shit.
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
		fmt.Fprintln(w, err)
		return
	}

	defer file.Close()

	out, err := os.Create("/tmp/uploadedfile")
	if err != nil {
		fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
		return
	}

	defer out.Close()

	// write the content from POST to the file
	_, err = io.Copy(out, file)
	if err != nil {
		fmt.Fprintln(w, err)
	}

	fmt.Fprintf(w, "File uploaded successfully : ")
	fmt.Fprintf(w, header.Filename)
}
