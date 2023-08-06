package repository

import (
	"douyin/util"
	"sync"

	"gorm.io/gorm"
)

type User struct {
	Id              int64  `gorm:"column:id"`
	Name            string `gorm:"column:name"`
	Password        string `gorm:"column:password"`
	FollowCount     int64  `gorm:"column:follow_count"`
	FollowerCount   int64  `gorm:"column:follower_count"`
	Avatar          string `gorm:"column:avatar"`
	BackgroundImage string `gorm:"column:background_image"`
	Signature       string `gorm:"column:signature"`
	TotalFavorited  int64  `gorm:"column:total_favorited"`
	WorkCount       int64  `gorm:"column:work_count"`
	FavoriteCount   int64  `gorm:"column:favorite_count"`
}

func (User) TableName() string {
	return "user"
}

type UserDao struct {
}

var userDao *UserDao
var userOnce sync.Once

func NewUserDaoInstance() *UserDao {
	userOnce.Do(
		func() {
			userDao = &UserDao{}
		})
	return userDao
}

func (*UserDao) QueryUserById(id int64) (*User, error) {
	var user User
	err := db.Where("id = ?", id).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		util.Logger.Error("find user by id err:" + err.Error())
		return nil, err
	}
	return &user, nil
}

func (*UserDao) QueryUserByName(name string) (*User, error) {
	var user User
	err := db.Where("Name = ?", name).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		util.Logger.Error("find user by Name err:" + err.Error())
		return nil, err
	}
	return &user, nil
}

func (*UserDao) MQueryUserByIds(ids []int64) ([]*User, error) {
	var users []*User
	err := db.Where("id in (?)", ids).Find(&users).Error
	if err != nil {
		util.Logger.Error("batch find user by ids err:" + err.Error())
		return nil, err
	}
	return users, nil
}

func (*UserDao) CreateUser(name string, password string) (error) {
	user := &User{
		Name: name,
		Password: password,
	}
	if err := db.Create(user).Error; err != nil {
		util.Logger.Error("insert user err:" + err.Error())
		return err
	}
	return nil
}


func (*UserDao) FavoriteCountUpdateById(id int64, favoriteCount int64) (error) {
	var user User
	err := db.Model(&user).Where("id = ?", id).Update("favorite_count", favoriteCount).Error
	if err != nil {
		util.Logger.Error("favorite count update by id err:" + err.Error())
		return err
	}
	return nil
}

func (*UserDao) TotalFavoriteUpdateById(id int64, totalFavorite int64) (error) {
	var user User
	err := db.Model(&user).Where("id = ?", id).Update("total_favorited", totalFavorite).Error
	if err != nil {
		util.Logger.Error("total favorite update by id err:" + err.Error())
		return err
	}
	return nil
}

func (*UserDao) WorkCountUpdateById(id int64, workCount int64) (error) {
	var user User
	err := db.Model(&user).Where("id = ?", id).Update("work_count", workCount).Error
	if err != nil {
		util.Logger.Error("work count update by id err:" + err.Error())
		return err
	}
	return nil
}

func (*UserDao) FollowCountUpdateById(id int64, followCount int64) (error) {
	var user User
	err := db.Model(&user).Where("id = ?", id).Update("follow_count", followCount).Error
	if err != nil {
		util.Logger.Error("follow count update by id err:" + err.Error())
		return err
	}
	return nil
}

func (*UserDao) FollowerCountUpdateById(id int64, followerCount int64) (error) {
	var user User
	err := db.Model(&user).Where("id = ?", id).Update("follower_count", followerCount).Error
	if err != nil {
		util.Logger.Error("follower count update by id err:" + err.Error())
		return err
	}
	return nil
}