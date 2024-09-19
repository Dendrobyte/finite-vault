package main

import (
	"fmt"
	"net/http"

	"github.com/Dendrobyte/finite_vault/auth"
	"github.com/Dendrobyte/finite_vault/db"
	"github.com/Dendrobyte/finite_vault/vault"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
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
 * [x] 1. Set up the OAuth stuff as mirrored in the JS side of things. Auth middleware
 * [x] 1.5. Make sure the JWT token is sent back, and also stored client-side for future requests the frontend sends
 * [ ] 2. Get the login data working, which means pulling data from mongo and sending it back with JWT token information
 * [ ] 3. Incrementing daily value when user logs in based on timestamp (this is beyond the JS stuff)
 * [ ] 4. Full transaction support via mongo, etc. (the actual meaty part, not sure if it'll take longer than auth)
 */

func main() {
	port := 5001
	fmt.Printf("-+- Server starting on port %d... -+-\n", port)

	db.InitMongoDB()

	router := chi.NewRouter()

	// Middleware (TODO: Authentication)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"}, //[]string{"http://localhost:5173", "http://10.0.0.*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	router.Get("/health", health)

	/* Login and Authentication */
	router.Post("/login/{service}", auth.LoginByService)

	router.Post("/validateToken", auth.ValidateJWT) // File under "/auth" if we make more auth related functions, e.g. wrapper for future access

	router.Get("/generateToken", auth.TestCreateJWT)

	/* Finite Vault Feature Routes */
	router.Get("/vaultBalance", vault.GetUserVaultBalance)
	// TODO: Make a route to update all data so the frontend can hit that if it "knows" enough time has elapsed

	router.Get("/getUserTransactions", vault.GetUserTransactions)

	router.Post("/newTransaction", vault.PostNewUserTransaction) // TODO: Put?

	/* Generic */
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// TODO: Permanent redirect for the base route?
		http.Redirect(w, r, "/health", http.StatusTemporaryRedirect)
	})

	fmt.Println("-+- Server started -+-")
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}
