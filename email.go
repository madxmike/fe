package fe

import (
	"strings"
)

type ValidationError string

func (e ValidationError) Error() string {
	return string(e)
}

var (
	ListIdEmpty           = ValidationError("listId cannot be empty")
	EmailAddressEmpty     = ValidationError("email address cannot empty")
	EmailAddressMalformed = ValidationError("email address is malformed")
)

type ListId string

func NewListId(s string) (ListId, error) {
	if s == "" {
		return "", ListIdEmpty
	}

	return ListId(s), nil
}

func (id *ListId) UnmarshalJSON(b []byte) error {
	new, err := NewListId(string(b))
	if err != nil {
		return err
	}

	*id = new
	return nil
}

type EmailAddress string

func NewEmailAddress(s string) (EmailAddress, error) {
	if s == "" {
		return "", EmailAddressEmpty
	}

	if !strings.Contains(s, "@") {
		return "", EmailAddressMalformed
	}

	return EmailAddress(s), nil
}

func (ea *EmailAddress) UnmarshalJSON(b []byte) error {
	new, err := NewEmailAddress(string(b))
	if err != nil {
		return err
	}

	*ea = new
	return nil
}
