package handlers

import (
	"github.com/gin-gonic/gin"
	"sloan.com/service/internal/event"
	eventrepository "sloan.com/service/internal/event-repository"
)

type Store struct {
	repoManager *eventrepository.RepositoryManager
}

type StoreGetResourceResponse struct {
	Events []*event.Event `json: "events"`
}

func (r *Store) GetResource(c *gin.Context) {
	name := c.Param("name")
	// should really clean this...but yolo
	repo, err := r.repoManager.GetRepo("")
	if (err) != nil {
		// handle not found
		switch err {
		case eventrepository.ErrRepoNotFound:
			c.JSON(404, nil)
			break
		default:
			c.AbortWithError(500, err)
			break
		}
	}
	evts, err := (*repo).GetResource(name)
	if err != nil {
		// handle
		switch err {
		case eventrepository.ErrResourceNotFound:
			c.AbortWithError(404, err)
			return
		default:
			c.AbortWithError(500, err)
			return
		}
	}
	c.JSON(200, &StoreGetResourceResponse{
		Events: evts,
	})
}
