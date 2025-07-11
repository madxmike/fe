package httpd

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/madxmike/fe"
)

func WriteError(w http.ResponseWriter, err error) {
	var validationError fe.ValidationError
	if errors.As(err, &validationError) {
		err = fmt.Errorf("validation failed: %w", validationError)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	return
}
