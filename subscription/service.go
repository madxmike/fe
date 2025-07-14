package subscription

import (
	"fmt"

	"github.com/madxmike/fe/valid"
)

type ListStorage interface {
	SaveSubscriberToList(listId valid.ListId, subscriberEmail valid.EmailAddress) error
	RemoveSubscriberToList(listId valid.ListId, subscriberEmail valid.EmailAddress) error
}

type Service struct {
	ListStorage ListStorage
}

func (s *Service) SubscribeToList(listId valid.ListId, subscriberEmail valid.EmailAddress) error {
	err := s.ListStorage.SaveSubscriberToList(listId, subscriberEmail)
	if err != nil {
		return fmt.Errorf("could not subscribe %s to list %s: %w", subscriberEmail, listId, err)
	}
	return nil
}

func (s *Service) UnsubscribeFromList(listId valid.ListId, subscriberEmail valid.EmailAddress) error {
	err := s.ListStorage.RemoveSubscriberToList(listId, subscriberEmail)
	if err != nil {
		return fmt.Errorf("could not unsubscribe %s to list %s: %w", subscriberEmail, listId, err)
	}
	return nil
}
