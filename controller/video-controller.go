package controller

import (
	"gilab.com/kmohd7121/golang-gin-poc/entity"
	"github.com/gin-gonic/gin"
	"gitlab.com/kmohd7121/golang-gin-poc/entity"
	"gitlab.com/pragmaticreviews/golang-gin-poc/service"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}
func (c *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.ShouldBindJSON(&video)
	c.service.Save(video)
	return video
}
