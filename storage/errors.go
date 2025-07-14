package storage

import "errors"

var (
	ErrorListDoesNotExist                  = errors.New("list does not exist")
	ErrorAlreadySubscribedToList           = errors.New("already subscribed to list")
	ErrorNoSubscriberExistsForEmailAddress = errors.New("no subscriber exists for the email address provides")
)
