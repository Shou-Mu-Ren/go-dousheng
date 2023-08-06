package service

import (
	"douyin/repository"
)

func IsFavorite(userId int64, videoId int64) (bool, error) {
	return repository.NewRedisDaoInstance().IsFavorite(userId, videoId)
}

func IsFollow(userId int64, toUserId int64) (bool, error) {
	return repository.NewRedisDaoInstance().IsFollow(userId, toUserId)
}
