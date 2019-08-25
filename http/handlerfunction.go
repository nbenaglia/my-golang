package main

import (
	"fmt"
	"net/http"
)

func bip(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bip!")
}

func bop(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bop!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/bip", bip)
	http.HandleFunc("/bop", bop)

	server.ListenAndServe()
}
