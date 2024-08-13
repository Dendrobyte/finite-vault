package auth

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

/*
 * Functions to handle login authentication, namely OAuth callback functions and generation of web tokens
 */

func LoginByService(w http.ResponseWriter, r *http.Request) {
	service := chi.URLParam(r, "service")
	token := r.URL.Query().Get("token")
	w.Write([]byte(fmt.Sprintf("Kachow! Got request from %s with token %s\n", service, token)))
}
