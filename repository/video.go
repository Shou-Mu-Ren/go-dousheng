package repository

import (
	"douyin/util"
	"time"

	"gorm.io/gorm"
	"sync"
)

type Video struct {
	Id            int64     `gorm:"column:id"`
	UserId        int64     `gorm:"column:user_id"`
	PlayUrl       string    `gorm:"column:play_url"`
	CoverUrl      string    `gorm:"column:cover_url"`
	FavoriteCount int64     `gorm:"column:favorite_count"`
	CommentCount  int64     `gorm:"column:comment_count"`
	Title         string    `gorm:"column:title"`
	CreateTime    time.Time `gorm:"column:create_time"`
}

func (Video) TableName() string {
	return "video"
}

type VideoDao struct {
}

var videoDao *VideoDao
var videoOnce sync.Once

func NewVideoDaoInstance() *VideoDao {
	videoOnce.Do(
		func() {
			videoDao = &VideoDao{}
		})
	return videoDao
}

func (*VideoDao) VideoUpload(userId int64, palyUrl string, coverUrl string, title string) error {
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	createTime, _ := time.ParseInLocation("2006-01-02 15:04:05", nowTime, time.Local)
	video := &Video{
		UserId:   userId,
		PlayUrl:  palyUrl,
		CoverUrl: coverUrl,
		Title:    title,
		CreateTime: createTime,
	}
	if err := db.Create(video).Error; err != nil {
		util.Logger.Error("insert video err:" + err.Error())
		return err
	}
	return nil
}

func (*VideoDao) QueryVideoById(id int64) (*Video, error) {
	var video Video
	err := db.Where("id = ?", id).Find(&video).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		util.Logger.Error("find video by id err:" + err.Error())
		return nil, err
	}
	return &video, nil
}

func (*VideoDao) MQueryVideoByUserId(userId int64) ([]*Video, error) {
	var videos []*Video
	err := db.Where("user_id = ? ", userId).Find(&videos).Error
	if err != nil {
		util.Logger.Error("batch find videos by user_id err:" + err.Error())
		return nil, err
	}
	return videos, nil
}

func (*VideoDao) MQueryVideosByIds(ids []int64) ([]*Video, error) {
	var videos []*Video
	err := db.Where("id in (?)", ids).Find(&videos).Error
	if err != nil {
		util.Logger.Error("batch find video by ids err:" + err.Error())
		return nil, err
	}
	return videos, nil
}

func (*VideoDao) MQueryVideoBeforeLastTime(lastTime string) ([]*Video, error) {
	var videos []*Video
	err := db.Where("create_time < ? ", lastTime).Order("create_time desc").Limit(30).Find(&videos).Error
	if err != nil {
		util.Logger.Error("batch find videos by last_time err:" + err.Error())
		return nil, err
	}
	return videos, nil
}

func (*VideoDao) FavoriteCountUpdateById(id int64, favoriteCount int64) (error) {
	var video Video
	err := db.Model(&video).Where("id = ?", id).Update("favorite_count", favoriteCount).Error
	if err != nil {
		util.Logger.Error("favorite count update by id err:" + err.Error())
		return err
	}
	return nil
}

func (*VideoDao) CommentCountUpdateById(id int64, commentCount int64) (error) {
	var video Video
	err := db.Model(&video).Where("id = ?", id).Update("comment_count", commentCount).Error
	if err != nil {
		util.Logger.Error("favorite count update by id err:" + err.Error())
		return err
	}
	return nil
}