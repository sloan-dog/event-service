package handlers

import (
	"github.com/gin-gonic/gin"
	eventrepository "sloan.com/service/internal/event-repository"
)

type Store struct {
	repoManager *eventrepository.RepositoryManager
}

func (r *Store) GetResource(c *gin.Context) {
	name := c.Param("name")
	// should really clean this...but yolo
	repo, err := r.repoManager.GetRepo("")
	if (err) != nil {
		// handle not found
		switch err {
		case eventrepository.ErrRepoNotFound:
			c.AbortWithStatus(404)
			break
		default:
			c.AbortWithError(500, err)
			break
		}
	}
	resource, err := (*repo).GetResource(name)
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
	cType := resource.GetContentType()
	data := resource.GetData()
	c.Data(200, cType, data)
}
