package httpd

import (
	"encoding/json"
	"net/http"

	"github.com/madxmike/fe/subscription"
	"github.com/madxmike/fe/valid"
)

type SubscriptionHandler struct {
	SubscriptionService subscription.Service
}

type SubscribeRequest struct {
	SubscriberEmail valid.EmailAddress `json:"emailAddress"`
}

func (h *SubscriptionHandler) Subscribe(w http.ResponseWriter, r *http.Request) {
	listId, err := ListIdURLParam(r)
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

	err = h.SubscriptionService.SubscribeToList(listId, request.SubscriberEmail)
	if err != nil {
		WriteError(w, err)
		return
	}
}

func (h *SubscriptionHandler) Unsubscribe(w http.ResponseWriter, r *http.Request) {
	listId, err := ListIdURLParam(r)
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

	err = h.SubscriptionService.UnsubscribeFromList(listId, request.SubscriberEmail)
	if err != nil {
		WriteError(w, err)
		return
	}
}
