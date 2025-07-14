package httpd

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/madxmike/fe/valid"
)

func WriteError(w http.ResponseWriter, err error) {
	var validationError valid.ValidationError
	if errors.As(err, &validationError) {
		err = fmt.Errorf("validation failed: %w", validationError)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	slog.Error(err.Error())
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func ListIdURLParam(r *http.Request) (valid.ListId, error) {
	raw, err := valid.NewNonEmptyString(chi.URLParam(r, "listId"))
	if err != nil {
		return "", err
	}
	listId, err := valid.NewListId(raw)
	if err != nil {
		return "", err
	}

	return listId, nil
}
