package main

import (
	"os"
	"net/http"

	"log"

	pb "github.com/Masterminds/go-in-practice/chapter10/userpb"
	"github.com/golang/protobuf/proto"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(res http.ResponseWriter, req *http.Request) {
	u := &pb.User{
		Name:  proto.String("Inigo Montoya"),
		Id:    proto.Int32(1234),
		Email: proto.String("inigo@montoya.example.com"),
	}

	body, err := proto.Marshal(u)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/x-protobuf")
	res.Write(body)

	f := log.Ldate | log.Lshortfile
	log.SetOutput(os.Stdout)
	log.SetFlags(f)
	log.Println("A message has been sent.")
}
