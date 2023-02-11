// init: $ go mod init github.com/USERNAME/go-blockchain
// dep: $ go get github.com/gorilla/mux, $ go install github.com/gorilla/mux
// run : $ go run main.go

package main

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World!")
}

func router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	return r
}

func main() {
	router := router()
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:9100",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
