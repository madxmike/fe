package httpd

import (
	"encoding/json"
	"net/http"

	"github.com/madxmike/fe/subscription"
	"github.com/madxmike/fe/valid"
)

type SubscribedMailingList struct {
	ID              valid.ID             `json:"id"`
	PublicationName valid.NonEmptyString `json:"publicationName"`
}

type SubscriptionHandler struct {
	SubscriptionService subscription.Service
}

type SubscribeRequest struct {
	SubscriberEmail valid.EmailAddress `json:"emailAddress"`
}

func (h *SubscriptionHandler) Subscribe(w http.ResponseWriter, r *http.Request) {
	id, err := ListIDURLParam(r)
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

	err = h.SubscriptionService.SubscribeToList(r.Context(), id, request.SubscriberEmail)
	if err != nil {
		WriteError(w, err)
		return
	}
}

func (h *SubscriptionHandler) Unsubscribe(w http.ResponseWriter, r *http.Request) {
	id, err := ListIDURLParam(r)
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

	err = h.SubscriptionService.UnsubscribeFromList(r.Context(), id, request.SubscriberEmail)
	if err != nil {
		WriteError(w, err)
		return
	}
}

type ListResponse struct {
	Lists []SubscribedMailingList
}

func (h *SubscriptionHandler) List(w http.ResponseWriter, r *http.Request) {
	subscriberId, err := SubscriberIDURLParam(r)
	if err != nil {
		WriteError(w, err)
		return
	}

	lists, err := h.SubscriptionService.GetSubscribedLists(r.Context(), subscriberId)
	if err != nil {
		WriteError(w, err)
		return
	}

	subscribedLists := make([]SubscribedMailingList, 0, len(lists))
	for _, v := range lists {
		subscribedLists = append(subscribedLists, SubscribedMailingList{
			ID:              v.ID,
			PublicationName: v.Publication.Name,
		})
	}

	response := ListResponse{
		Lists: subscribedLists,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		WriteError(w, err)
		return
	}
}
