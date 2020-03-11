package handlers

import (
	"github.com/gin-gonic/gin"
	"sloan.com/service/internal/middleware"
)

type Stats struct {
}

func (s *Stats) Stats(ctx *gin.Context) {
	ctx.JSON(200, middleware.GetStats())
}
