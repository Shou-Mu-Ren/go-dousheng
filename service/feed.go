package service

import (
	"douyin/repository"
)

func FindVideosBeforeLastTime(lastTime string) ([]*repository.Video, error) {
	return repository.NewVideoDaoInstance().MQueryVideoBeforeLastTime(lastTime)
}

func FindUserById(id int64) (*repository.User, error) {
	return repository.NewUserDaoInstance().QueryUserById(id)
}
