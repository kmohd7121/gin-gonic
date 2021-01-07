package service

import "gitlab.com/pragmaticreviews/golang-gin-poc/entity"

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}
