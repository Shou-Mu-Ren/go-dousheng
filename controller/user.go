package controller

import (
	"douyin/service"
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	if exist, err := service.UserExist(username); exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}

	token, _ := MakeToken(username, password)

	if err := service.UserRegister(username, password); err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	
	if user, err := service.FindUserByName(username); err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	if user, err := service.FindUserByNameAndPassword(username, password); err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	} else {
		token, _ := MakeToken(username, password)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")

	if user, err := ValidToken(token); err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	} else {
		isFollow, _ := service.IsFollow(user.Id, user.Id)
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     User{
				Id: user.Id,
				Name: user.Name,
				FollowCount: user.FollowCount,
				FollowerCount: user.FollowerCount,
				IsFollow: isFollow,
				Avatar: user.Avatar,
				BackgroundImage: user.BackgroundImage,
				Signature: user.Signature,
				TotalFavorited: user.TotalFavorited,
				WorkCount: user.WorkCount,
				FavoriteCount: user.FavoriteCount,
			},
		})
	}
}
