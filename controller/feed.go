package controller

import (
	"douyin/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	lastTimeString := c.Query("latest_time")
	lastTimeInt, _ := strconv.ParseInt(lastTimeString, 10, 64)
	lastTime := time.Unix(lastTimeInt, 0).Format("2006-01-02 15:04:05")

	vs, err := service.FindVideosBeforeLastTime(lastTime)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}

	vds, err := service.FindVideosBeforeLastTime(time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}

	token := c.Query("token")
	usr, err := ValidToken(token)
	if err != nil {
		var nextTime time.Time
		if len(vs) != 0 {
			videos := make([]Video, len(vs))

			for index, video := range vs {
				u, err := service.FindUserById(video.UserId)
				if err != nil {
					c.JSON(http.StatusOK, UserResponse{
						Response: Response{StatusCode: 1, StatusMsg: err.Error()},
					})
					return
				}
				user := User{
					Id:              u.Id,
					Name:            u.Name,
					FollowCount:     u.FollowCount,
					FollowerCount:   u.FollowerCount,
					IsFollow:        false,
					Avatar:          u.Avatar,
					BackgroundImage: u.BackgroundImage,
					Signature:       u.Signature,
					TotalFavorited:  u.TotalFavorited,
					WorkCount:       u.WorkCount,
					FavoriteCount:   u.FavoriteCount,
				}
				videos[index] = Video{
					Id:            video.Id,
					Author:        user,
					PlayUrl:       video.PlayUrl,
					CoverUrl:      video.CoverUrl,
					FavoriteCount: video.FavoriteCount,
					CommentCount:  video.CommentCount,
					IsFavorite:    false,
					Title:         video.Title,
				}
				nextTime = video.CreateTime
			}
			c.JSON(http.StatusOK, FeedResponse{
				Response:  Response{StatusCode: 0},
				VideoList: videos,
				NextTime:  nextTime.Unix(),
			})

		} else {
			videos := make([]Video, len(vds))

			for index, video := range vds {
				u, err := service.FindUserById(video.UserId)
				if err != nil {
					c.JSON(http.StatusOK, UserResponse{
						Response: Response{StatusCode: 1, StatusMsg: err.Error()},
					})
					return
				}
				user := User{
					Id:              u.Id,
					Name:            u.Name,
					FollowCount:     u.FollowCount,
					FollowerCount:   u.FollowerCount,
					IsFollow:        false,
					Avatar:          u.Avatar,
					BackgroundImage: u.BackgroundImage,
					Signature:       u.Signature,
					TotalFavorited:  u.TotalFavorited,
					WorkCount:       u.WorkCount,
					FavoriteCount:   u.FavoriteCount,
				}
				videos[index] = Video{
					Id:            video.Id,
					Author:        user,
					PlayUrl:       video.PlayUrl,
					CoverUrl:      video.CoverUrl,
					FavoriteCount: video.FavoriteCount,
					CommentCount:  video.CommentCount,
					IsFavorite:    false,
					Title:         video.Title,
				}
				nextTime = video.CreateTime
			}
			c.JSON(http.StatusOK, FeedResponse{
				Response:  Response{StatusCode: 0},
				VideoList: videos,
				NextTime:  nextTime.Unix(),
			})
		}
		return
	}

	var nextTime time.Time
	if len(vs) != 0 {
		videos := make([]Video, len(vs))

		for index, video := range vs {
			u, err := service.FindUserById(video.UserId)
			if err != nil {
				c.JSON(http.StatusOK, UserResponse{
					Response: Response{StatusCode: 1, StatusMsg: err.Error()},
				})
				return
			}
			isFollow, _ := service.IsFollow(usr.Id, u.Id)
			user := User{
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

			isFavorite, _ := service.IsFavorite(usr.Id, video.Id)
			videos[index] = Video{
				Id:            video.Id,
				Author:        user,
				PlayUrl:       video.PlayUrl,
				CoverUrl:      video.CoverUrl,
				FavoriteCount: video.FavoriteCount,
				CommentCount:  video.CommentCount,
				IsFavorite:    isFavorite,
				Title:         video.Title,
			}
			nextTime = video.CreateTime
		}
		c.JSON(http.StatusOK, FeedResponse{
			Response:  Response{StatusCode: 0},
			VideoList: videos,
			NextTime:  nextTime.Unix(),
		})

	} else {
		videos := make([]Video, len(vds))

		for index, video := range vds {
			u, err := service.FindUserById(video.UserId)
			if err != nil {
				c.JSON(http.StatusOK, UserResponse{
					Response: Response{StatusCode: 1, StatusMsg: err.Error()},
				})
				return
			}
			isFollow, _ := service.IsFollow(usr.Id, u.Id)
			user := User{
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

			isFavorite, _ := service.IsFavorite(usr.Id, video.Id)
			videos[index] = Video{
				Id:            video.Id,
				Author:        user,
				PlayUrl:       video.PlayUrl,
				CoverUrl:      video.CoverUrl,
				FavoriteCount: video.FavoriteCount,
				CommentCount:  video.CommentCount,
				IsFavorite:    isFavorite,
				Title:         video.Title,
			}
			nextTime = video.CreateTime
		}
		c.JSON(http.StatusOK, FeedResponse{
			Response:  Response{StatusCode: 0},
			VideoList: videos,
			NextTime:  nextTime.Unix(),
		})
	}
}
