package service

import "douyin/repository"

func VideoUpload(user *repository.User, palyUrl string, coverUrl string, title string) error {
	if err := repository.NewVideoDaoInstance().VideoUpload(user.Id, palyUrl, coverUrl, title); err != nil{
		return err
	}
	user.WorkCount += 1
	if err := repository.NewUserDaoInstance().WorkCountUpdateById(user.Id, user.WorkCount); err != nil{
		return err
	}
	return nil
}


func FindVideosByUserId(userId int64) ([]*repository.Video, error){
	return repository.NewVideoDaoInstance().MQueryVideoByUserId(userId)
}

