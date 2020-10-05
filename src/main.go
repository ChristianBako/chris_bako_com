package main

import (
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
