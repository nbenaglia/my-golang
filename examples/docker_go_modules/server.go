package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", rootHandler)

	fmt.Println("Listening on : 9000")
	http.ListenAndServe(":9000", r)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Docker test server running!")

}
