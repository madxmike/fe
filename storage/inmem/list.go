package inmem

import (
	"context"

	"github.com/google/uuid"
	"github.com/madxmike/fe/list"
	"github.com/madxmike/fe/storage"
	"github.com/madxmike/fe/valid"
)

func (s *Storage) SaveList(ctx context.Context, mailingList list.MailingList) (list.MailingList, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	uuid, err := uuid.NewV7()
	if err != nil {
		return list.MailingList{}, err
	}

	mailingList.ID = valid.ID(uuid)
	s.lists = append(s.lists, mailingList)

	s.subscriptions[mailingList.ID] = make([]valid.ID, 0)

	return mailingList, nil
}

func (s *Storage) ReadList(ctx context.Context, id valid.ID) (list.MailingList, error) {
	for _, mailingList := range s.lists {
		if mailingList.ID == id {
			return mailingList, nil
		}
	}

	return list.MailingList{}, storage.ErrorListDoesNotExist
}
