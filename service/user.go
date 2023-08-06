package service

import (
	"douyin/repository"
	"errors"
	// "fmt"
)

func UserExist(name string)(bool, error){
	user, err := repository.NewUserDaoInstance().QueryUserByName(name)
	if err != nil{
		return true, err
	}
	if user.Id == 0 {
		return false, nil
	}
	return true, errors.New("用户已存在")
}


func UserRegister(name string, password string)(error){
	if err := repository.NewUserDaoInstance().CreateUser(name, password) ; err != nil{
		return errors.New("用户注册失败")
	} else {
		return nil
	}
}

func FindUserByName(name string)(*repository.User, error){
	if user, err := repository.NewUserDaoInstance().QueryUserByName(name) ; err != nil{
		return nil, err
	} else {
		if user.Id == 0 {
			return nil, errors.New("用户不存在")
		} else {
			return user, nil
		}
	}
}

func FindUserByNameAndPassword(name string, password string)(*repository.User, error){
	if user, err := repository.NewUserDaoInstance().QueryUserByName(name) ; err != nil{
		return nil, err
	} else {
		if user.Id == 0 {
			return nil, errors.New("用户不存在")
		} else if user.Password != password {
			return nil, errors.New("密码错误")
		} else {
			return user, nil
		}
	}
}