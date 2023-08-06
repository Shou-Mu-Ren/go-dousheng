package controller

import (
	"douyin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

type FriendUserListResponse struct {
	Response
	UserList []FriendUser `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	user, err := ValidToken(token)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	actionType := c.Query("action_type")
	toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	toUser, _ := service.FindUserById(toUserId)

	if err := service.RelationAction(user, toUser, actionType); err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	token := c.Query("token")
	usr, err := ValidToken(token)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	us, err := service.FollowList(usr.Id)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	users := make([]User, len(us))
	for index, user := range us {
		users[index] = User{
			Id:              user.Id,
			Name:            user.Name,
			FollowCount:     user.FollowCount,
			FollowerCount:   user.FollowerCount,
			IsFollow:        true,
			Avatar:          user.Avatar,
			BackgroundImage: user.BackgroundImage,
			Signature:       user.Signature,
			TotalFavorited:  user.TotalFavorited,
			WorkCount:       user.WorkCount,
			FavoriteCount:   user.FavoriteCount,
		}
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: users,
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	token := c.Query("token")
	usr, err := ValidToken(token)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	us, err := service.FollowerList(usr.Id)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	users := make([]User, len(us))
	for index, user := range us {
		isFollow, _ := service.IsFollow(usr.Id, user.Id)
		users[index] = User{
			Id:              user.Id,
			Name:            user.Name,
			FollowCount:     user.FollowCount,
			FollowerCount:   user.FollowerCount,
			IsFollow:        isFollow,
			Avatar:          user.Avatar,
			BackgroundImage: user.BackgroundImage,
			Signature:       user.Signature,
			TotalFavorited:  user.TotalFavorited,
			WorkCount:       user.WorkCount,
			FavoriteCount:   user.FavoriteCount,
		}
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: users,
	})
}

// FriendList all users have same friend list
func FriendList(c *gin.Context) {
	token := c.Query("token")
	usr, err := ValidToken(token)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}

	us, err := service.FriendList(usr.Id)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}

	users := make([]FriendUser, len(us))
	for index, user := range us {
		isFollow, _ := service.IsFollow(usr.Id, user.Id)

		u := User{
			Id:              user.Id,
			Name:            user.Name,
			FollowCount:     user.FollowCount,
			FollowerCount:   user.FollowerCount,
			IsFollow:        isFollow,
			Avatar:          user.Avatar,
			BackgroundImage: user.BackgroundImage,
			Signature:       user.Signature,
			TotalFavorited:  user.TotalFavorited,
			WorkCount:       user.WorkCount,
			FavoriteCount:   user.FavoriteCount,
		}

		mes, _ := service.GetNewMessageByUserIdAndToUserId(usr.Id, u.Id)

		var i int
		if mes.FromUserId == usr.Id && mes.ToUserId == u.Id {
			i = 1
		} else {
			i = 0
		}

		users[index] = FriendUser{
			User:    u,
			Message: mes.Content,
			MsgType: i,
		}
	}

	c.JSON(http.StatusOK, FriendUserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: users,
	})
}
