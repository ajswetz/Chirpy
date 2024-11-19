package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiConfig) metricsHandler(resWriter http.ResponseWriter, _ *http.Request) {

	resWriter.Header().Add("Content-Type:", "text/html")
	resWriter.WriteHeader(200)

	// Get current file server hits from server state, convert to string, build full body text
	pageHits := cfg.fileserverHits.Load()

	htmlText := fmt.Sprintf(`
<html>
  <body>
    <h1>Welcome, Chirpy Admin</h1>
    <p>Chirpy has been visited %d times!</p>
  </body>
</html>
	`, pageHits)

	// Write response text using .Write() method
	resWriter.Write([]byte(htmlText))

}
