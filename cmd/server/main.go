package main

import (
	"context"
	"flag"
	"github.com/whosonfirst/go-http-whosonfirst-data"
	"github.com/whosonfirst/go-reader"
	"log"
	"net/http"
)

func main() {

	reader_uri := flag.String("reader-uri", "", "A valid whosonfirst/go-reader URI")
	server_uri := flag.String("server-uri", "localhost:8080", "The host address and port to listen for requests on")

	flag.Parse()

	ctx := context.Background()

	r, err := reader.NewReader(ctx, *reader_uri)

	if err != nil {
		log.Fatalf("Failed to create reader for '%s', %v", *server_uri, err)
	}

	data_handler := data.WhosOnFirstDataHandler(r)

	mux := http.NewServeMux()
	mux.Handle("/", data_handler)

	log.Printf("Listening for requests on %s\n", *server_uri)
	
	err = http.ListenAndServe(*server_uri, mux)

	if err != nil {
		log.Fatal("Failed to serve requests, %v", err)
	}
}
