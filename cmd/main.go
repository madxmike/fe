package main

import (
	"fmt"
	"net/http"

	"github.com/madxmike/fe/httpd"
	"github.com/madxmike/fe/list"
	"github.com/madxmike/fe/storage/inmem"
	"github.com/madxmike/fe/subscription"
)

func main() {
	fmt.Println("Hello, World!")

	inmemStorage := inmem.NewStorage()
	subscriptionService := subscription.Service{
		ListStore:       &inmemStorage,
		SubscriberStore: &inmemStorage,
	}

	listService := list.Service{
		ListStore: &inmemStorage,
	}

	mux := httpd.BuildRoutes(httpd.RouteHandlers{
		Subscription: httpd.SubscriptionHandler{
			SubscriptionService: subscriptionService,
		},
		List: httpd.ListHandler{
			ListService: listService,
		},
	})

	http.ListenAndServe(":3333", mux)
}
