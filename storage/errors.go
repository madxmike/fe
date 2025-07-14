package storage

import "errors"

var (
	ErrorListDoesNotExist        = errors.New("list does not exist")
	ErrorAlreadySubscribedToList = errors.New("already subscribed to list")
)
