package graph

import (
	"gituhub.com/TheDSCPL/gql/graph/model"
	"sync"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	usersMutex  sync.RWMutex
	usersByID   map[model.ID]*model.User
	usersByName map[string]*model.User

	todos []*model.Todo
}

func (r *Resolver) Init() {
	if r == nil {
		return
	}

	r.usersMutex.Lock()
	defer r.usersMutex.Unlock()

	if r.usersByID == nil {
		r.usersByID = make(map[model.ID]*model.User)
	}

	if r.usersByName == nil {
		r.usersByName = make(map[string]*model.User)
	}
}
