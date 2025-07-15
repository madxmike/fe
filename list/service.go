package list

import (
	"context"
	"fmt"

	"github.com/madxmike/fe/valid"
)

const (
	PublicationDefaultName = "New Publication"
)

type MailingList struct {
	ID           valid.ID
	Admin        valid.ID
	EmailAddress valid.EmailAddress
	Publication  Publication
}

type Publication struct {
	Name valid.NonEmptyString
}

type ListStore interface {
	SaveList(ctx context.Context, list MailingList) (MailingList, error)
	ReadList(ctx context.Context, listId valid.ID) (MailingList, error)
}

type Service struct {
	ListStore ListStore
}

func (s *Service) CreateList(ctx context.Context, adminID valid.ID, emailAddress valid.EmailAddress) (MailingList, error) {
	list := MailingList{
		Admin:        adminID,
		EmailAddress: emailAddress,
		Publication: Publication{
			Name: PublicationDefaultName,
		},
	}
	list, err := s.ListStore.SaveList(ctx, list)
	if err != nil {
		return MailingList{}, fmt.Errorf("failed to create list: %w", err)
	}

	return list, nil
}

func (s *Service) Info(ctx context.Context, listId valid.ID) (MailingList, error) {
	list, err := s.ListStore.ReadList(ctx, listId)
	if err != nil {
		return MailingList{}, fmt.Errorf("failed to get list info: %w", err)
	}

	return list, nil
}
