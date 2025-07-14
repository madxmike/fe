package inmem

import (
	"sync"

	"github.com/madxmike/fe/list"
	"github.com/madxmike/fe/subscription"
	"github.com/madxmike/fe/valid"
)

type Storage struct {
	mu *sync.Mutex

	lists         []list.MailingList
	subscribers   []subscription.Subscriber
	subscriptions map[valid.ID][]valid.ID
}

func NewStorage() Storage {
	return Storage{
		mu: &sync.Mutex{},

		lists:         make([]list.MailingList, 0),
		subscribers:   make([]subscription.Subscriber, 0),
		subscriptions: make(map[valid.ID][]valid.ID),
	}
}
