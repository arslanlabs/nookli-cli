package auth

import (
	"github.com/go-chi/chi/v5"
)

// Router builds the /api/v1/auth sub-router
func Router() chi.Router {
	r := chi.NewRouter()
	r.Post("/signup", signUp)
	r.Post("/login", login)
	r.Post("/refresh", refresh)
	r.Post("/logout", logout)
	r.Post("/recover", recoverPassword)
	return r
}
