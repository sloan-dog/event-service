package eventrepository

import (
	"errors"

	"sloan.com/service/internal/event"
)

var (
	ErrResourceNotFound = errors.New("resource not found")
	ErrRepoNotFound     = errors.New("repo not found")
)

type EventRepository interface {
	GetResource(name string) ([]*event.Event, error)
	Name() string
}

type RepositoryManager struct {
	repositories map[string]*EventRepository
}

func NewRepositoryManager(repositories ...EventRepository) *RepositoryManager {
	r := RepositoryManager{
		repositories: map[string]*EventRepository{},
	}
	for _, val := range repositories {
		if val != nil {
			r.repositories[val.Name()] = &val
		}
	}
	return &r
}

func (rm *RepositoryManager) GetRepo(name string) (*EventRepository, error) {
	if name == "" {
		name = defaultRepo
	}
	if val, ok := rm.repositories[name]; ok {
		return val, nil
	}
	return nil, ErrRepoNotFound
}
