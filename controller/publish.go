package controller

import (
	"douyin/service"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}


var baseUrl = "http://192.168.1.103:8080/douyin/file/download?filename="

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")

	user, err := ValidToken(token);
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	if err := service.VideoUpload(user, baseUrl + finalName, baseUrl + "666.jpg", filename); err != nil{
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}


	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}


	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + "上传成功",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	token := c.Query("token")

	u, err := ValidToken(token);
	if  err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	isFollow, _ := service.IsFollow(u.Id, u.Id)
	user := User{
		Id: u.Id,
		Name: u.Name,
		FollowCount: u.FollowCount,
		FollowerCount: u.FollowerCount,
		IsFollow: isFollow,
		Avatar: u.Avatar,
		BackgroundImage: u.BackgroundImage,
		Signature: u.Signature,
		TotalFavorited: u.TotalFavorited,
		WorkCount: u.WorkCount,
		FavoriteCount: u.FavoriteCount,
	}

	vs, err := service.FindVideosByUserId(user.Id)
	if  err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}

	videos := make([]Video,len(vs))
	for index, video := range vs {
		isFavorite, _ :=  service.IsFavorite(user.Id, video.Id)
		videos[index] = Video{
			Id: video.Id,
			Author: user,
			PlayUrl: video.PlayUrl,
			CoverUrl: video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount: video.CommentCount,
			IsFavorite: isFavorite,
			Title: video.Title,
		}
	}

	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: videos,
	})
}
