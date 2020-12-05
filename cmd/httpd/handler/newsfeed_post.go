package handler

import (
	"gin-api/cmd/platform/newsfeed"
	"net/http"

	"github.com/gin-gonic/gin"
)

// declare the request
type newsfeedPostRequest struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

func NewsfeedPost(feed newsfeed.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqBody := newsfeedPostRequest{}
		c.Bind(&reqBody)
		item := newsfeed.Item{
			Title: reqBody.Title,
			Post:  reqBody.Post,
		}
		feed.Add(item)
		c.Status(http.StatusNoContent)
	}
}
