package service

import (
	"douyin/repository"
	"errors"
)



func MessageAction(userId int64, toUserId int64, actionType string, content string)(error){
	if actionType == "1"{
		return repository.NewMessageDaoInstance().CreateMessage(userId, toUserId, content)
	}
	return errors.New("请求错误");
}

func MessageChat(userId int64, toUserId int64, preMsgTime int64)([]*repository.Message, error){
	return repository.NewMessageDaoInstance().MQueryMessageByUserIdAndToUserId(userId, toUserId, preMsgTime)
}

