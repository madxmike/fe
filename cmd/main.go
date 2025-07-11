package main

import (
	"fmt"
	"net/http"

	"github.com/madxmike/fe/httpd"
)

func main() {
	fmt.Println("Hello, World!")

	mux := httpd.BuildRoutes(httpd.RouteHandlers{
		Subscription: httpd.SubscriptionHandler{},
	})

	http.ListenAndServe(":3333", mux)
}
