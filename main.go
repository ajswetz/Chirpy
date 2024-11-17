package main

import (
	"fmt"
	"net/http"
)

func main() {

	serveMux := http.NewServeMux()

	server := http.Server{
		Addr:    ":8080",
		Handler: serveMux,
	}

	fmt.Println("Starting Chirpy . . . ")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
	}

}
