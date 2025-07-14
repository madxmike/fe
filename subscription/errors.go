package subscription

import "errors"

var (
	ErrorNoSubscriberExistsForEmailAddress = errors.New("no subscriber exists for provided email address")
	ErrorNoSubscriberExistsForId           = errors.New("no subscriber exists for provided id")
)
