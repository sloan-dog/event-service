package handlers

import "github.com/gin-gonic/gin"

type Root struct {
}

type RootResponse struct {
	Ok   bool   `json:"ok"`
	Cake string `json:"cake"`
}

func (r *Root) Root(c *gin.Context) {
	c.JSON(200, &RootResponse{
		Ok:   true,
		Cake: "foo",
	})
}
