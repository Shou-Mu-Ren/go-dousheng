package controller

import (
	"douyin/service"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ChatResponse struct {
	Response
	MessageList []Message `json:"message_list"`
}

// MessageAction no practical effect, just check if token is valid
func MessageAction(c *gin.Context) {
	token := c.Query("token")
	user, err := ValidToken(token)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	content := c.Query("content")
	actionType := c.Query("action_type")
	if err := service.MessageAction(user.Id, toUserId, actionType, content); err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	}
}

// MessageChat all users have same follow list
func MessageChat(c *gin.Context) {
	token := c.Query("token")
	user, err := ValidToken(token)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}

	toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	preMsgTime, _ := strconv.ParseInt(c.Query("pre_msg_time"), 10, 64)

	fmt.Println(preMsgTime)

	ms, err := service.MessageChat(user.Id, toUserId, preMsgTime)
	messages := make([]Message, len(ms))
	for index, message := range ms {
		t, _ := time.ParseInLocation("2006-01-02 15:04:05", message.CreateTime.Format("2006-01-02 15:04:05"), time.UTC)
		messages[index] = Message{
			Id:         message.Id,
			ToUserId:   message.ToUserId,
			FromUserId: message.FromUserId,
			Content:    message.Content,
			CreateTime: t.Unix(),
		}
	}
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
	} else {
		c.JSON(http.StatusOK, ChatResponse{Response: Response{StatusCode: 0}, MessageList: messages})
	}
}
