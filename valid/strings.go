package valid

import (
	"encoding/json"
	"strings"

	"github.com/google/uuid"
)

var (
	StringIsNullOrEmpty         = ValidationError("value is null or empty")
	ListIdNotUUID               = ValidationError("list id is not a valid uuid value")
	EmailAddressMalformed       = ValidationError("email address is malformed")
	EmailAddressPrefixMalformed = ValidationError("email address prefix is malformed")
	EmailAddressDomainMalformed = ValidationError("email address domain is malformed")
)

type NonEmptyString string

func NewNonEmptyString(s string) (NonEmptyString, error) {
	if s == "" {
		return "", StringIsNullOrEmpty
	}

	return NonEmptyString(s), nil
}

func (s *NonEmptyString) UnmarshalJSON(b []byte) error {
	var raw string
	err := json.Unmarshal(b, &raw)
	if err != nil {
		return err
	}

	v, err := NewNonEmptyString(raw)
	if err != nil {
		return err
	}

	*s = v
	return nil
}

type EmailAddress NonEmptyString

func NewEmailAddressFromRaw(s NonEmptyString) (EmailAddress, error) {
	if !strings.Contains(string(s), "@") {
		return "", EmailAddressMalformed
	}

	return EmailAddress(s), nil
}

func NewEmailAddressFromParts(prefix NonEmptyString, domain NonEmptyString) (EmailAddress, error) {
	return EmailAddress(prefix + "@" + domain), nil
}

func (ea *EmailAddress) UnmarshalJSON(b []byte) error {
	var raw NonEmptyString
	err := json.Unmarshal(b, &raw)
	if err != nil {
		return err
	}

	v, err := NewEmailAddressFromRaw(raw)
	if err != nil {
		return err
	}

	*ea = v
	return nil
}

type ID uuid.UUID

func NewListId(s NonEmptyString) (ID, error) {
	uuid, err := uuid.Parse(string(s))
	if err != nil {
		return ID{}, ListIdNotUUID
	}

	return ID(uuid), nil
}

func (id *ID) UnmarshalJSON(b []byte) error {
	var raw NonEmptyString
	err := json.Unmarshal(b, &raw)
	if err != nil {
		return err
	}

	listId, err := NewListId(raw)
	if err != nil {
		return err
	}

	*id = listId
	return nil
}
