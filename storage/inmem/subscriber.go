package inmem

import (
	"context"
	"slices"

	"github.com/google/uuid"
	"github.com/madxmike/fe/list"
	"github.com/madxmike/fe/storage"
	"github.com/madxmike/fe/subscription"
	"github.com/madxmike/fe/valid"
)

func (s *Storage) SaveSubscriber(ctx context.Context, subscriber subscription.Subscriber) (subscription.Subscriber, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if subscriber.ID == valid.ID(uuid.Nil) {
		uuid, err := uuid.NewV7()
		if err != nil {
			return subscription.Subscriber{}, err
		}

		subscriber.ID = valid.ID(uuid)
	}

	s.subscribers = append(s.subscribers, subscriber)
	return subscriber, nil
}

func (s *Storage) SaveSubscriberToList(ctx context.Context, listId valid.ID, subscriberEmail valid.EmailAddress) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	subscriber, err := s.GetSubscriberFromEmailAddress(context.TODO(), subscriberEmail)
	if err != nil {
		return err
	}

	listSubscribers, ok := s.subscriptions[listId]
	if !ok {
		return storage.ErrorListDoesNotExist
	}

	if slices.Contains(listSubscribers, subscriber.ID) {
		return storage.ErrorAlreadySubscribedToList
	}

	s.subscriptions[listId] = append(listSubscribers, subscriber.ID)
	return nil
}

func (s *Storage) RemoveSubscriberToList(ctx context.Context, listId valid.ID, subscriberEmail valid.EmailAddress) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	subscriber, err := s.GetSubscriberFromEmailAddress(context.TODO(), subscriberEmail)
	if err != nil {
		return err
	}

	listSubscribers, ok := s.subscriptions[listId]
	if !ok {
		return storage.ErrorListDoesNotExist
	}

	idx := 0
	for i, v := range listSubscribers {
		if v == subscriber.ID {
			idx = i
			break
		}
	}

	s.subscriptions[listId] = append(listSubscribers[:idx], listSubscribers[idx+1:]...)
	return nil
}

func (s *Storage) GetSubscribedLists(ctx context.Context, subscriber subscription.Subscriber) ([]list.MailingList, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	subscribedLists := make([]list.MailingList, 0)
	for _, list := range s.lists {
		listSubscribers, ok := s.subscriptions[list.ID]
		if !ok {
			continue
		}

		for _, listSubscriber := range listSubscribers {
			if listSubscriber == subscriber.ID {
				subscribedLists = append(subscribedLists, list)
			}
		}
	}

	return subscribedLists, nil
}

func (s *Storage) GetSubscriberFromEmailAddress(ctx context.Context, emailAddress valid.EmailAddress) (subscription.Subscriber, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, subscriber := range s.subscribers {
		if subscriber.EmailAddress == emailAddress {
			return subscriber, nil
		}
	}

	return subscription.Subscriber{}, storage.ErrorNoSubscriberExistsForEmailAddress
}

func (s *Storage) GetSubscriberFromId(ctx context.Context, id valid.ID) (subscription.Subscriber, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, subscriber := range s.subscribers {
		if subscriber.ID == id {
			return subscriber, nil
		}
	}

	return subscription.Subscriber{}, storage.ErrorNoSubscriberExistsForEmailAddress
}
