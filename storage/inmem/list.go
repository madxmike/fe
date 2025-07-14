package inmem

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"slices"
	"sync"

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

	// TODO (Michael): Create more elaborate listid system than just hash , probably uuidv7 or something
	hash := md5.Sum([]byte(emailAddress))
	listId, err := valid.NewListId(valid.NonEmptyString(string(hex.EncodeToString(hash[:]))))
	if err != nil {
		return "", err
	}

	s.lists[listId] = make([]valid.EmailAddress, 0)

	return listId, nil
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
