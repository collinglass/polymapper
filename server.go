package main

import (
	// "encoding/json"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Starting Server")
	http.Handle("/", logHandler(http.FileServer(http.Dir("./app/"))))

	log.Println("Listening on 8000")
	panic(http.ListenAndServe(":8000", nil))
}

func logHandler(h http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, h)
}

/*
func jsonHandler(h http.Handler) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		h.ServeHTTP(rw, req)
	}
}
*/
