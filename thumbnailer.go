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
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", RootHandler)
	r.HandleFunc("/thumbnail", ThumbnailHandler)
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
	switch r.Method {
	case "GET":
		fmt.Fprint(w, "I only handle POSTs")
	case "POST":
		if size := r.FormValue("size"); size == "" {
			fmt.Fprint(w, "You've forgotten something")
		} else {
			fmt.Fprint(w, "Sizing to "+size)
		}
	}
}
