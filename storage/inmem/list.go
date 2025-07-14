package inmem

import (
	"errors"
	"slices"
	"sync"

	"github.com/google/uuid"
	"github.com/madxmike/fe/valid"
)

type ListStorage struct {
	mu *sync.Mutex

	lists map[valid.ListId][]valid.EmailAddress
}

func NewListStorage() ListStorage {
	return ListStorage{
		mu:    &sync.Mutex{},
		lists: make(map[valid.ListId][]valid.EmailAddress),
	}
}

func (s *ListStorage) CreateList(emailAddress valid.EmailAddress) (valid.ListId, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	uuid, err := uuid.NewV7()
	if err != nil {
		return valid.ListId{}, err
	}

	listID := valid.ListId(uuid)

	// TODO (Michael): Duplication error
	s.lists[listID] = make([]valid.EmailAddress, 0)

	return listID, nil
}

func (s *ListStorage) SaveSubscriberToList(listId valid.ListId, subscriberEmail valid.EmailAddress) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	list := s.lists[listId]
	if list == nil {
		return errors.New("list does not exist")
	}

	if slices.Contains(list, subscriberEmail) {
		return errors.New("subscriber is already subscribed")
	}

	s.lists[listId] = append(s.lists[listId], subscriberEmail)
	return nil
}

func (s *ListStorage) RemoveSubscriberToList(listId valid.ListId, subscriberEmail valid.EmailAddress) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	list := s.lists[listId]
	if list == nil {
		return errors.New("list does not exist")
	}

	idx := slices.Index(list, subscriberEmail)
	if idx == -1 {
		return nil
	}

	s.lists[listId] = append(list[:idx], list[idx+1:]...)
	return nil
}
