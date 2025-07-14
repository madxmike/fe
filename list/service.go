package list

import (
	"fmt"

	"github.com/madxmike/fe/valid"
)

type MailingList struct {
	ID           valid.ID
	EmailAddress valid.EmailAddress
}

type ListStorage interface {
	CreateList(emailAddress valid.EmailAddress) (valid.ID, error)
	ReadList(id valid.ID) (MailingList, error)
}

type Service struct {
	ListStorage ListStorage
}

func (s *Service) CreateList(emailAddress valid.EmailAddress) (valid.ID, error) {
	listId, err := s.ListStorage.CreateList(emailAddress)
	if err != nil {
		return valid.ID{}, fmt.Errorf("failed to create list: %w", err)
	}

	return listId, nil
}

func (s *Service) Info(id valid.ID) (MailingList, error) {
	list, err := s.ListStorage.ReadList(id)
	if err != nil {
		return MailingList{}, fmt.Errorf("failed to get list info: %w", err)
	}

	return list, nil
}
