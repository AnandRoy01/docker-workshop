package main

import (
	"net/http"
	"github.com/Pallinder/sillyname-go"
	"log"
)

const (
	PORT = ":9000"
)


func main () {

	http.HandleFunc("/hello", func(w http.ResponseWriter, r* http.Request){
		_, err := w.Write([]byte("{\"message\":\"Hi, "+sillyname.GenerateStupidName()+"\"}"))
		if(err != nil){
			panic(err)
		}
	});

	log.Println("Listening server on "+PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
