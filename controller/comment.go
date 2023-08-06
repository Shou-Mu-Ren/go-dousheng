package controller

import (
	"douyin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	user, err := ValidToken(token)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	actionType := c.Query("action_type")
	comment := c.Query("comment_text")
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	commentId, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)

	com, err := service.CommentAction(videoId, user.Id, comment, actionType, commentId)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	if actionType == "1" {
		isFollow, _ := service.IsFollow(user.Id, user.Id)
		c.JSON(http.StatusOK, CommentActionResponse{Response: Response{StatusCode: 0},
			Comment: Comment{
				Id: com.Id,
				User: User{
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
				},
				Content:    com.Content,
				CreateDate: com.CreateDate,
			}})
		return
	}
	c.JSON(http.StatusOK, Response{StatusCode: 0})
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	token := c.Query("token")
	usr, err := ValidToken(token)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}

	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)

	coms, err := service.CommentList(videoId)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}

	comments := make([]Comment, len(coms))
	for index, comment := range coms {
		u, _ := service.FindUserById(comment.UserId)
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

		comments[index] = Comment{
			Id:         comment.Id,
			User:       user,
			Content:    comment.Content,
			CreateDate: comment.CreateDate,
		}
	}

	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0},
		CommentList: comments,
	})
}
