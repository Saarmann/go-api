package handler

import (
	"gin-api/cmd/platform/newsfeed"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewsfeedGet(feed newsfeed.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		res := feed.GetAll()
		c.JSON(http.StatusOK, res)
	}
}
