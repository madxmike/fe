package httpd

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type RouteHandlers struct {
	Subscription SubscriptionHandler
}

func BuildRoutes(routeHandlers RouteHandlers) http.Handler {
	r := chi.NewRouter()

	r.Route("/subscriptions", func(r chi.Router) {
		r.Route("/{listId}", func(r chi.Router) {
			r.Post("/subscribe", routeHandlers.Subscription.Subscribe)
		})
	})
	return r
}
