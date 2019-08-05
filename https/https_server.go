package main

import (
	"log"
	"net/http"
	"os"
)

func foo(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is a http example server.\n"))
}

func main() {
	log.SetOutput(os.Stdout)
	log.Println("HTTPS server starting...")
	http.HandleFunc("/", foo)
	http.ListenAndServeTLS(":1044", "server.crt", "server.key", nil)
}
