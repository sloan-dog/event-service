package handlers

import (
	"net/http"

	eventrepository "sloan.com/service/internal/event-repository"
	"sloan.com/service/internal/platform/web"
)

func API() http.Handler {

	a := web.NewApp()

	a.Router.Use()

	fileRepo, err := eventrepository.NewFileRepository("/tmp/foo/repos")
	if err != nil {
		// do something
	}
	repoMan := eventrepository.NewRepositoryManager(fileRepo)

	root := Root{}
	health := Health{}
	store := Store{
		repoManager: repoMan,
	}

	a.Handle("GET", "/v1", root.Root)
	a.Handle("GET", "/v1/health-check", health.Check)
	a.Handle("GET", "/v1/store/:name", store.GetResource)
	return a
}
