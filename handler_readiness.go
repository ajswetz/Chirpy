package main

import "net/http"

func readinessHandler(resWriter http.ResponseWriter, _ *http.Request) {

	resWriter.Header().Add("Content-Type:", "text/plain; charset=utf-8")
	resWriter.WriteHeader(200)
	resWriter.Write([]byte("OK"))

}
