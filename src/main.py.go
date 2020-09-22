package main

import (
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./"))
	http.Handle("/", fileServer)
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}


func index(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("heyo"))
}
