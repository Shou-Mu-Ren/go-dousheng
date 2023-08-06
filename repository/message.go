package repository

import (
	// "gorm.io/gorm"
	"douyin/util"
	"sync"
	"time"
)

type Message struct {
	Id         int64     `gorm:"column:id"`
	ToUserId   int64     `gorm:"column:to_user_id"`
	FromUserId int64     `gorm:"column:from_user_id"`
	Content    string    `gorm:"column:content"`
	CreateTime time.Time `gorm:"column:create_time"`
}

func (Message) TableName() string {
	return "message"
}

type MessageDao struct {
}

var messageDao *MessageDao
var messageOnce sync.Once

func NewMessageDaoInstance() *MessageDao {
	messageOnce.Do(
		func() {
			messageDao = &MessageDao{}
		})
	return messageDao
}

func (*MessageDao) CreateMessage(userId int64, toUserId int64, content string) error {
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	createTime, _ := time.ParseInLocation("2006-01-02 15:04:05", nowTime, time.Local)
	message := &Message{
		ToUserId:   toUserId,
		FromUserId: userId,
		Content:    content,
		CreateTime: createTime,
	}
	if err := db.Create(message).Error; err != nil {
		util.Logger.Error("insert message err:" + err.Error())
		return err
	}
	return nil
}

func (*MessageDao) MQueryMessageByUserIdAndToUserId(userId int64, toUserId int64, createTime int64) ([]*Message, error) {
	var messages []*Message
	ts := time.Unix(createTime, 0).Format("2006-01-02 15:04:05")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", ts, time.Local)
	err := db.Where("from_user_id = ? AND to_user_id = ? AND create_time > ? ", userId, toUserId, t).
	Or("from_user_id = ? AND to_user_id = ? AND create_time > ? ", toUserId, userId, t).Find(&messages).Error
	if err != nil {
		util.Logger.Error("batch find messages by user_id an to_user_from err:" + err.Error())
		return nil, err
	}
	return messages, nil
}

func (*MessageDao) QueryMessageByUserIdAndToUserId(userId int64, toUserId int64) (*Message, error) {
	var messages []*Message
	err := db.Where("from_user_id = ? AND to_user_id = ? ", userId, toUserId).
	Or("from_user_id = ? AND to_user_id = ? ", toUserId, userId).Order("create_time desc").Find(&messages).Error
	if err != nil {
		util.Logger.Error("batch find messages by user_id an to_user_from err:" + err.Error())
		return nil, err
	}
	if len(messages) == 0 {
		nowTime := time.Now().Format("2006-01-02 15:04:05")
		createTime, _ := time.ParseInLocation("2006-01-02 15:04:05", nowTime, time.Local)
		return &Message{Id: 0, ToUserId: 0, FromUserId: 0, Content: "无消息", CreateTime: createTime}, nil
	}
	return messages[0], nil
}
