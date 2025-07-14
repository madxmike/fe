package subscription

import (
	"context"
	"fmt"

	"github.com/madxmike/fe/list"
	"github.com/madxmike/fe/valid"
)

type ListStore interface {
	SaveSubscriberToList(ctx context.Context, listId valid.ID, subscriber Subscriber) error
	RemoveSubscriberFromList(ctx context.Context, listId valid.ID, subscriber Subscriber) error
}

type SubscriberStore interface {
	SaveSubscriber(ctx context.Context, subscriber Subscriber) (Subscriber, error)
	GetSubscribedLists(ctx context.Context, subscriber Subscriber) ([]list.MailingList, error)
	GetSubscriberFromEmailAddress(ctx context.Context, subscriberEmail valid.EmailAddress) (Subscriber, error)
	GetSubscriberFromId(ctx context.Context, subscriberId valid.ID) (Subscriber, error)
}

type Subscriber struct {
	ID           valid.ID
	EmailAddress valid.EmailAddress
}

type Service struct {
	ListStore       ListStore
	SubscriberStore SubscriberStore
}

func (s *Service) SubscribeToList(ctx context.Context, listId valid.ID, subscriberEmail valid.EmailAddress) error {
	subscriber, err := s.SubscriberStore.GetSubscriberFromEmailAddress(ctx, subscriberEmail)
	if err != nil {
		return fmt.Errorf("could not subscribe to list: %w", err)
	}

	err = s.ListStore.SaveSubscriberToList(ctx, listId, subscriber)
	if err != nil {
		return fmt.Errorf("could not subscribe to list: %w", err)
	}
	return nil
}

func (s *Service) UnsubscribeFromList(ctx context.Context, listId valid.ID, subscriberEmail valid.EmailAddress) error {
	subscriber, err := s.SubscriberStore.GetSubscriberFromEmailAddress(ctx, subscriberEmail)
	if err != nil {
		return fmt.Errorf("could not unsubscribe from list: %w", err)
	}

	err = s.ListStore.RemoveSubscriberFromList(ctx, listId, subscriber)
	if err != nil {
		return fmt.Errorf("could not unsubscribe from list: %w", err)
	}
	return nil
}

func (s *Service) GetSubscribedLists(ctx context.Context, subscriberId valid.ID) ([]list.MailingList, error) {
	subscriber, err := s.SubscriberStore.GetSubscriberFromId(ctx, subscriberId)
	if err != nil {
		return nil, fmt.Errorf("could not get subscribed lists: %w", err)
	}

	lists, err := s.SubscriberStore.GetSubscribedLists(ctx, subscriber)
	if err != nil {
		return nil, fmt.Errorf("could not get subscribed lists: %w", err)
	}

	return lists, nil
}
