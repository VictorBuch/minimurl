package application

import (
	"github.com/VictorBuch/link-shortener/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Route("/shorten-link", loadShortLinkRoutes)

	return router
}

func loadShortLinkRoutes(router chi.Router) {
	shortLinkHandler := &handler.ShortLink{}

	router.Post("/", shortLinkHandler.Create)
	router.Get("/", shortLinkHandler.List)
	router.Get("/{id}", shortLinkHandler.GetByID)
	router.Put("/{id}", shortLinkHandler.UpdateByID)
	router.Delete("/{id}", shortLinkHandler.DeleteByID)
}
