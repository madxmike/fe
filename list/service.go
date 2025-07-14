package list

import (
	"fmt"

	"github.com/madxmike/fe/valid"
)

type ListStorage interface {
	CreateList(emailAddress valid.EmailAddress) (valid.ListId, error)
}

type Service struct {
	ListStorage ListStorage
}

func (s *Service) CreateList(emailAddress valid.EmailAddress) (valid.ListId, error) {
	listId, err := s.ListStorage.CreateList(emailAddress)
	if err != nil {
		return valid.ListId{}, fmt.Errorf("failed to create list: %w", err)
	}

	return listId, nil
}
