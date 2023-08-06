package repository

import (
	// "gorm.io/gorm"
	"douyin/util"
	"sync"

	"gorm.io/gorm"
)

type Comment struct {
	Id         int64  `gorm:"column:id"`
	UserId     int64  `gorm:"column:user_id"`
	VideoId    int64  `gorm:"column:video_id"`
	Content    string `gorm:"column:content"`
	CreateDate string `gorm:"column:create_date"`
}

func (Comment) TableName() string {
	return "comment"
}

type CommentDao struct {
}

var commentDao *CommentDao
var commentOnce sync.Once

func NewCommentDaoInstance() *CommentDao {
	commentOnce.Do(
		func() {
			commentDao = &CommentDao{}
		})
	return commentDao
}

func (*CommentDao) CreateComment(videoId int64, userId int64, content string, createDate string) error {
	comment := &Comment{
		UserId:     userId,
		VideoId:    videoId,
		Content:    content,
		CreateDate: createDate,
	}
	if err := db.Create(comment).Error; err != nil {
		util.Logger.Error("insert comment err:" + err.Error())
		return err
	}
	return nil
}

func (*CommentDao) QueryCommentByAll(videoId int64, userId int64, content string, createDate string) (*Comment, error) {
	var comment Comment
	err := db.Where("video_id = ? AND user_id = ? AND content = ? AND create_date = ? ", videoId, userId, content, createDate).Find(&comment).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		util.Logger.Error("find comment by all err:" + err.Error())
		return nil, err
	}
	return &comment, nil
}

func (*CommentDao) DeleteCommentById(id int64) error {
	return db.Delete(Comment{}, "id = ? ", id).Error
}


func (*CommentDao) MQueryCommentByVideoId(videoId int64) ([]*Comment, error) {
	var comments []*Comment
	err := db.Where("video_id = ? ", videoId).Find(&comments).Error
	if err != nil {
		util.Logger.Error("batch find comments by video_id err:" + err.Error())
		return nil, err
	}
	return comments, nil
}