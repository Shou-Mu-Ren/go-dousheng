package service

import (
	"douyin/repository"
	"strconv"
)

func RelationAction(user *repository.User, toUser *repository.User, actionType string) error {
	if actionType == "1" {
		if err := repository.NewRedisDaoInstance().AddFollow(user.Id, toUser.Id); err != nil {
			return err
		}
		if err := repository.NewRedisDaoInstance().AddFollower(user.Id, toUser.Id); err != nil {
			return err
		}
		user.FollowCount += 1
		if err := repository.NewUserDaoInstance().FollowCountUpdateById(user.Id, user.FollowCount); err != nil {
			return err
		}
		toUser.FollowerCount += 1
		if err := repository.NewUserDaoInstance().FollowerCountUpdateById(toUser.Id, toUser.FollowerCount); err != nil {
			return err
		}
	} else {
		if err := repository.NewRedisDaoInstance().RemFollow(user.Id, toUser.Id); err != nil {
			return err
		}
		if err := repository.NewRedisDaoInstance().RemFollower(user.Id, toUser.Id); err != nil {
			return err
		}
		user.FollowCount -= 1
		if err := repository.NewUserDaoInstance().FollowCountUpdateById(user.Id, user.FollowCount); err != nil {
			return err
		}
		toUser.FollowerCount -= 1
		if err := repository.NewUserDaoInstance().FollowerCountUpdateById(toUser.Id, toUser.FollowerCount); err != nil {
			return err
		}
	}
	return nil
} 


func FollowList(userId int64)([]*repository.User, error){
	userIds, err := repository.NewRedisDaoInstance().GetFollowUserIds(userId)
	if err != nil {
		return nil, err
	}
	ids := make([]int64, len(userIds))

	for index, id := range userIds{
		ids[index], _ = strconv.ParseInt(id, 10, 64)
	}

	return repository.NewUserDaoInstance().MQueryUserByIds(ids)
}

func FollowerList(userId int64)([]*repository.User, error){
	userIds, err := repository.NewRedisDaoInstance().GetFollowerUserIds(userId)
	if err != nil {
		return nil, err
	}
	ids := make([]int64, len(userIds))

	for index, id := range userIds{
		ids[index], _ = strconv.ParseInt(id, 10, 64)
	}

	return repository.NewUserDaoInstance().MQueryUserByIds(ids)
}

func FriendList(userId int64)([]*repository.User, error){
	userIds, err := repository.NewRedisDaoInstance().GetFriendUserIds(userId)
	if err != nil {
		return nil, err
	}
	ids := make([]int64, len(userIds))

	for index, id := range userIds{
		ids[index], _ = strconv.ParseInt(id, 10, 64)
	}

	return repository.NewUserDaoInstance().MQueryUserByIds(ids)
}

func GetNewMessageByUserIdAndToUserId(userId int64, toUserId int64)(*repository.Message, error){
	return repository.NewMessageDaoInstance().QueryMessageByUserIdAndToUserId(userId, toUserId,)
}