package router

import (
	"github.com/ckcherry23/folioboard-backend/internal/routes"
	"github.com/go-chi/chi"
)

func Setup() chi.Router {
	r := chi.NewRouter()
	setUpRoutes(r)
	return r
}

func setUpRoutes(r chi.Router) {
	r.Group(routes.GetRoutes())
}
