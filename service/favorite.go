package service

import (
	"douyin/repository"
	"strconv"
)

func FavoriteAction(user *repository.User, videoId int64, actionType string)(error){
	video, _ := repository.NewVideoDaoInstance().QueryVideoById(videoId)
	if actionType == "1"{
		if err := repository.NewRedisDaoInstance().AddFavorite(user.Id,videoId); err != nil {
			return err
		}
		video.FavoriteCount += 1
		if err := repository.NewVideoDaoInstance().FavoriteCountUpdateById(videoId, video.FavoriteCount); err != nil {
			return err
		}
		user.FavoriteCount += 1
		if err := repository.NewUserDaoInstance().FavoriteCountUpdateById(user.Id, user.FavoriteCount); err != nil {
			return err
		}
		u, _ := repository.NewUserDaoInstance().QueryUserById(video.UserId)
		u.TotalFavorited += 1
		if err := repository.NewUserDaoInstance().TotalFavoriteUpdateById(u.Id, u.TotalFavorited); err != nil {
			return err
		}
	}else{
		if err := repository.NewRedisDaoInstance().RemFavorite(user.Id,videoId); err != nil {
			return err
		}
		video.FavoriteCount -= 1
		if err := repository.NewVideoDaoInstance().FavoriteCountUpdateById(videoId, video.FavoriteCount); err != nil {
			return err
		}
		user.FavoriteCount -= 1	
		if err := repository.NewUserDaoInstance().FavoriteCountUpdateById(user.Id, video.FavoriteCount); err != nil {
			return err
		}
		u, _ := repository.NewUserDaoInstance().QueryUserById(video.UserId)
		u.TotalFavorited -= 1	
		if err := repository.NewUserDaoInstance().TotalFavoriteUpdateById(u.Id, u.TotalFavorited); err != nil {
			return err
		}
	}
	return nil
}


func FavoriteList(userId int64)([]*repository.Video, error){
	videoIds, err := repository.NewRedisDaoInstance().GetFavoriteVideoIds(userId)
	if err != nil {
		return nil, err
	}
	ids := make([]int64, len(videoIds))

	for index, id := range videoIds{
		ids[index], _ = strconv.ParseInt(id, 10, 64)
	}

	return repository.NewVideoDaoInstance().MQueryVideosByIds(ids)

}