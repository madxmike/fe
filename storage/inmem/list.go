package inmem

import (
	"slices"
	"sync"

	"github.com/google/uuid"
	"github.com/madxmike/fe/list"
	"github.com/madxmike/fe/storage"
	"github.com/madxmike/fe/valid"
)

type ListStorage struct {
	mu *sync.Mutex

	lists         []list.MailingList
	subscriptions map[valid.ID][]valid.EmailAddress
}

func NewListStorage() ListStorage {
	return ListStorage{
		mu:            &sync.Mutex{},
		lists:         make([]list.MailingList, 0),
		subscriptions: make(map[valid.ID][]valid.EmailAddress),
	}
}

func (s *ListStorage) CreateList(emailAddress valid.EmailAddress) (valid.ID, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	uuid, err := uuid.NewV7()
	if err != nil {
		return valid.ID{}, err
	}

	listID := valid.ID(uuid)

	s.lists = append(s.lists, list.MailingList{
		ID:           listID,
		EmailAddress: emailAddress,
	})

	s.subscriptions[listID] = make([]valid.EmailAddress, 0)

	return listID, nil
}

func (s *ListStorage) SaveSubscriberToList(id valid.ID, subscriberEmail valid.EmailAddress) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	listSubscribers, ok := s.subscriptions[id]
	if !ok {
		return storage.ErrorListDoesNotExist
	}

	if slices.Contains(listSubscribers, subscriberEmail) {
		return storage.ErrorAlreadySubscribedToList
	}

	s.subscriptions[id] = append(listSubscribers, subscriberEmail)
	return nil
}

func (s *ListStorage) RemoveSubscriberToList(id valid.ID, subscriberEmail valid.EmailAddress) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	listSubscribers, ok := s.subscriptions[id]
	if !ok {
		return storage.ErrorListDoesNotExist
	}

	idx := slices.Index(listSubscribers, subscriberEmail)
	if idx == -1 {
		return nil
	}

	s.subscriptions[id] = append(listSubscribers[:idx], listSubscribers[idx+1:]...)
	return nil
}

func (s *ListStorage) ReadList(id valid.ID) (list.MailingList, error) {
	for _, list := range s.lists {
		if list.ID == id {
			return list, nil
		}
	}

	return list.MailingList{}, storage.ErrorListDoesNotExist
}
