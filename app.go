package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/changelog/", http.StripPrefix("/changelog/", http.FileServer(http.Dir("./changelog"))))
	http.Handle("/progress/", http.StripPrefix("/progress/", http.FileServer(http.Dir("./progress"))))
	http.Handle("/", http.FileServer(http.Dir("./index")))

	addr := ":3000"
	log.Printf("Server starting on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
