package httpd

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type RouteHandlers struct {
	Subscription SubscriptionHandler
	List         ListHandler
}

func BuildRoutes(routeHandlers RouteHandlers) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/subscriptions", func(r chi.Router) {
		r.Route("/{listId}", func(r chi.Router) {
			r.Post("/unsubscribe", routeHandlers.Subscription.Unsubscribe)
		})
	})

	r.Route("/lists", func(r chi.Router) {
		r.Post("/register", routeHandlers.List.Register)
		r.Route("/{listId}", func(r chi.Router) {
			r.Get("/info", routeHandlers.List.Info)
			r.Post("/subscribe", routeHandlers.Subscription.Subscribe)
		})
	})
	return r
}
