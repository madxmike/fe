package valid

import (
	"encoding/json"
	"strings"
)

var (
	StringIsNullOrEmpty         = ValidationError("value is null or empty")
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

type ListId NonEmptyString

func NewListId(s NonEmptyString) (ListId, error) {
	return ListId(s), nil
}

func (id *ListId) UnmarshalJSON(b []byte) error {
	var raw NonEmptyString
	err := json.Unmarshal(b, &raw)
	if err != nil {
		return err
	}

	v, err := NewListId(raw)
	if err != nil {
		return err
	}

	*id = v
	return nil
}
