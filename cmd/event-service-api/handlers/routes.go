package handlers

import (
	"net/http"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"sloan.com/service/internal/constants"
	eventrepository "sloan.com/service/internal/event-repository"
	"sloan.com/service/internal/middleware"
	"sloan.com/service/internal/platform/web"
)

func API() http.Handler {

	a := web.NewApp(logger.SetLogger())

	rawFileRepo, err := eventrepository.NewRawFileRepository(constants.TmpDir)
	if err != nil {
		// do something
	}
	repoMan := eventrepository.NewRepositoryManager(rawFileRepo)

	root := Root{}
	health := Health{}
	store := Store{
		repoManager: repoMan,
	}
	stats := Stats{}

	// default route
	a.Router.NoRoute(middleware.StatsMiddleware(), func(c *gin.Context) {
		c.JSON(404, gin.H{
			"message": "page not found",
		})
	})

	a.Handle("GET", "/v1", middleware.StatsMiddleware(), root.Root)
	a.Handle("GET", "/v1/health-check", middleware.StatsMiddleware(), health.Check)
	a.Handle("GET", "/v1/store/:name", middleware.StatsMiddleware(), store.GetResource)
	a.Handle("GET", "/v1/stats", stats.Stats) // no stats for you
	return a
}
