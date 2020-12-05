package main

import (
	"gin-api/cmd/httpd/handler"
	"gin-api/cmd/platform/newsfeed"

	"github.com/gin-gonic/gin"
)

func main() {
	feed := newsfeed.New()
	r := gin.Default()

	r.GET("/ping", handler.PingGet())
	r.GET("/news", handler.NewsfeedGet(feed))
	r.POST("/news", handler.NewsfeedPost(feed))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
