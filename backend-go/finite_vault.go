package main

import (
	"fmt"
	"net/http"
)

type HealthResponse struct {
	Message string
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "おはようございます!"}`)) // avoiding proper marshaling for simplicity
}

/*
 * Shit that needs to get done
 * [x] 0. Choose a framework (let's just use net/http since fasthttp is overkill and not for this)
 * [ ] 1. Set up the OAuth stuff as mirrored in the JS side of things. Auth guards??
 * [ ] 2. Get the login data working, which means pulling data from mongo and sending it back
 * [ ] 3. Incrementing daily value when user logs in based on timestamp (this is beyond the JS stuff)
 * [ ] 4. Full transaction support via mongo, etc. (the actual meaty part, not sure if it'll take longer than auth)
 */

func main() {
	port := 5000
	fmt.Printf("Server starting on port %d...\n", port)

	// server := &http.Server{
	// 	Addr:    fmt.Sprintf(":%d", port),
	// 	Handler: nil, // TODO: httprouter package?
	// }

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Permanent redirection, no timeout
		// http.Redirect(w, r, "/health", http.StatusPermanentRedirect)
	})

	http.HandleFunc("/health", health)

	http.ListenAndServe(fmt.Sprintf("localhost:%d", port), nil)
}
