package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {

	port := flag.Int("port", 8080, "Port to listen on (default: 8080)")
	dir := flag.String("dir", ".", "Directory to server from (default: .)")
	flag.Parse()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), http.FileServer(http.Dir(*dir))))
}
