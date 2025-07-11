package httpd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/madxmike/fe"
)

type SubscriptionHandler struct {
}

type SubscribeRequest struct {
	SubscriberEmail fe.EmailAddress `json:"emailAddress"`
}

func (h *SubscriptionHandler) Subscribe(w http.ResponseWriter, r *http.Request) {
	listId, err := fe.NewListId(chi.URLParam(r, "listId"))
	if err != nil {
		WriteError(w, err)
		return
	}

	var request SubscribeRequest
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		WriteError(w, err)
		return
	}

	fmt.Printf("%s", listId)
	fmt.Printf("%s", request.SubscriberEmail)
}
