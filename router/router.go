package router

import (
	"douyin/controller/video"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	routerGroup := r.Group("douyin/")
	{
		routerGroup.GET("/feed/", video.FeedVideosHandler)
	}
	return r
}
