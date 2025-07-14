package httpd

import (
	"encoding/json"
	"net/http"

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
	ListId valid.ListId `json:"listId"`
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

	listId, err := h.ListService.CreateList(emailAddress)
	if err != nil {
		WriteError(w, err)
		return
	}

	response := RegisterResponse{
		ListId: listId,
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		WriteError(w, err)
		return
	}
}
