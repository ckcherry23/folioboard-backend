package routes

import (
	"encoding/json"
	"net/http"

	"github.com/ckcherry23/folioboard-backend/internal/handlers/users"
	"github.com/go-chi/chi"
)

func GetRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/users", func(w http.ResponseWriter, req *http.Request) {
			response, _ := users.HandleList(w, req)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		})
	}
}
