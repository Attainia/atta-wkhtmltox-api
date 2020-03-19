package main

import (
	"atta-wkhtmltox-api/internal"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	config := internal.GetConfig()
	wkhtmltoxView := internal.WkhtmltoxView{}

	// Ensure the working directory exists
	_ = os.MkdirAll(config.WorkDir, os.ModePerm)

	http.HandleFunc("/", wkhtmltoxView.Handle)

	log.Printf("Listening on %s:%d", config.Host, config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", config.Host, config.Port), logRequest(http.DefaultServeMux)))
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
