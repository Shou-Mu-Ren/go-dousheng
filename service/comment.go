package service

import (
	"douyin/repository"
	"time"
)

func CommentAction(videoId int64, userId int64, content string, actionType string, commentId int64) (*repository.Comment, error) {

	createDate := time.Now().Format("01-02")
	video, err := repository.NewVideoDaoInstance().QueryVideoById(videoId)
	if err != nil {
		return nil, err
	}
	if actionType == "1" {
		if err := repository.NewCommentDaoInstance().CreateComment(videoId, userId, content, createDate); err != nil {
			return nil, err
		}
		video.CommentCount += 1
		repository.NewVideoDaoInstance().CommentCountUpdateById(video.Id, video.CommentCount)
		return repository.NewCommentDaoInstance().QueryCommentByAll(videoId, userId, content, createDate)

	} else {
		if err := repository.NewCommentDaoInstance().DeleteCommentById(commentId); err != nil {

			return nil, err
		}
		video.CommentCount -= 1
		repository.NewVideoDaoInstance().CommentCountUpdateById(video.Id, video.CommentCount)
		return nil, nil
	}
}

func CommentList(videoId int64) ([]*repository.Comment, error) {
	return repository.NewCommentDaoInstance().MQueryCommentByVideoId(videoId)
}
