package main

import (
	"fmt"
	"net/http"
)

type BipHandler struct{}

func (h *BipHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bip!")
}

type BopHandler struct{}

func (h *BopHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bop!")
}

func main() {
	bip := BipHandler{}
	bop := BopHandler{}

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.Handle("/bip", &bip)
	http.Handle("/bop", &bop)

	server.ListenAndServe()
}
