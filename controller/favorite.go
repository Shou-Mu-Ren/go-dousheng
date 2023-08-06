package controller

import (
	"douyin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	user, err := ValidToken(token)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}

	videoId := c.Query("video_id")
	actionType := c.Query("action_type")
	id, _ := strconv.ParseInt(videoId, 10, 64)

	if err := service.FavoriteAction(user,id,actionType); err !=nil{
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
	}else{
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	token := c.Query("token")

	user, err := ValidToken(token)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}

	vs, er := service.FavoriteList(user.Id) 
	if er != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: er.Error()},
		})
		return
	}

	videos := make([]Video, len(vs))
	for index, video := range vs {
		u, err := service.FindUserById(video.UserId)
		if err != nil {
			c.JSON(http.StatusOK, UserResponse{
				Response: Response{StatusCode: 1, StatusMsg: err.Error()},
			})
			return
		}
		isFollow, _ := service.IsFollow(user.Id, u.Id)
		usr := User{
			Id:              u.Id,
			Name:            u.Name,
			FollowCount:     u.FollowCount,
			FollowerCount:   u.FollowerCount,
			IsFollow:        isFollow,
			Avatar:          u.Avatar,
			BackgroundImage: u.BackgroundImage,
			Signature:       u.Signature,
			TotalFavorited:  u.TotalFavorited,
			WorkCount:       u.WorkCount,
			FavoriteCount:   u.FavoriteCount,
		}
		videos[index] = Video{
			Id:            video.Id,
			Author:        usr,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    true,
			Title:         video.Title,
		}
	}



	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: videos,
	})
}
