package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Health struct {
}

func (r *Health) Check(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusNoContent)
}
