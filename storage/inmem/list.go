package inmem

import (
	"github.com/google/uuid"
	"github.com/madxmike/fe/list"
	"github.com/madxmike/fe/storage"
	"github.com/madxmike/fe/valid"
)

func (s *Storage) CreateList(emailAddress valid.EmailAddress) (valid.ID, error) {
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

	s.subscriptions[listID] = make([]valid.ID, 0)

	return listID, nil
}

func (s *Storage) ReadList(id valid.ID) (list.MailingList, error) {
	for _, list := range s.lists {
		if list.ID == id {
			return list, nil
		}
	}

	return list.MailingList{}, storage.ErrorListDoesNotExist
}
