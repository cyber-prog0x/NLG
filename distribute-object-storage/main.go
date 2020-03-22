package main

import (
	"github.com/cyber-prog0x/NLG/distribute-object-storage/objects"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
