package httpd

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/madxmike/fe/list"
	"github.com/madxmike/fe/valid"
)

type ListHandler struct {
	ListService list.Service
}

type RegisterRequest struct {
	Prefix valid.NonEmptyString `json:"prefix"`
	Domain valid.NonEmptyString `json:"domain"`
}

type RegisterResponse struct {
	ListId valid.ID `json:"listId"`
}

func (h *ListHandler) Register(w http.ResponseWriter, r *http.Request) {
	var request RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		WriteError(w, err)
		return
	}

	emailAddress, err := valid.NewEmailAddressFromParts(request.Prefix, request.Domain)
	if err != nil {
		WriteError(w, err)
		return
	}

	// TODO (Michael): We want this to be given from the request. Maybe we need to look this up in the service level instead?
	adminUUID, err := uuid.NewV7()
	if err != nil {
		WriteError(w, err)
		return
	}
	adminId := valid.ID(adminUUID)

	list, err := h.ListService.CreateList(r.Context(), adminId, emailAddress)
	if err != nil {
		WriteError(w, err)
		return
	}

	response := RegisterResponse{
		ListId: list.ID,
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		WriteError(w, err)
		return
	}
}

type InfoResponse struct {
	ID           valid.ID           `json:"id"`
	EmailAddress valid.EmailAddress `json:"emailAddress"`
}

func (h *ListHandler) Info(w http.ResponseWriter, r *http.Request) {
	id, err := ListIDURLParam(r)
	if err != nil {
		WriteError(w, err)
		return
	}

	list, err := h.ListService.Info(r.Context(), id)
	if err != nil {
		WriteError(w, err)
		return
	}

	response := InfoResponse{
		ID:           list.ID,
		EmailAddress: list.EmailAddress,
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		WriteError(w, err)
		return
	}
}
