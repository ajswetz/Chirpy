package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	// Use the http.NewServeMux's .Handle() method to add a handler for the root path (/).
	// Use a standard http.FileServer as the handler
	// Use http.Dir to convert a filepath (in our case a dot: . which indicates the current directory) to a directory for the http.FileServer.
	mux.Handle("/app/", http.StripPrefix("/app/", http.FileServer(http.Dir("."))))
	mux.HandleFunc("/healthz", readinessHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Starting Chirpy . . . ")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
	}

}

func readinessHandler(resWriter http.ResponseWriter, _ *http.Request) {

	resWriter.Header().Add("Content-Type:", "text/plain; charset=utf-8")
	resWriter.WriteHeader(200)
	resWriter.Write([]byte("OK"))

}
