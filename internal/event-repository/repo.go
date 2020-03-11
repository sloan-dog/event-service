package eventrepository

import (
	"errors"
)

var (
	ErrResourceNotFound = errors.New("resource not found")
	ErrRepoNotFound     = errors.New("repo not found")
)

var (
	defaultRepo = "defaultRepo"
)

type Resource interface {
	GetContentType() string
	GetData() []byte
}

type Repository interface {
	GetResource(name string) (Resource, error)
	Name() string
}

type RepositoryManager struct {
	repositories map[string]*Repository
}

func NewRepositoryManager(repositories ...Repository) *RepositoryManager {
	r := RepositoryManager{
		repositories: map[string]*Repository{},
	}
	for _, val := range repositories {
		if val != nil {
			r.repositories[val.Name()] = &val
		}
	}
	return &r
}

func (rm *RepositoryManager) GetRepo(name string) (*Repository, error) {
	if name == "" {
		name = defaultRepo
	}
	if val, ok := rm.repositories[name]; ok {
		return val, nil
	}
	return nil, ErrRepoNotFound
}
