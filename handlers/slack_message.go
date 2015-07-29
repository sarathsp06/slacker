package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sarath.sp06/slacker/es"
	"github.com/sarath.sp06/slacker/types"
)

//SlackMessagePostHandler handles the outgoing webhook from slack
//to mock the request do this
//  curl -XPOST -H "Content-Type:application/x-www-form-urlencoded"  http://127.0.0.1:8080/slackmessage --data "token=sasasasaegxsdh&team_id=T03PTHYK9&team_domain=robohome&service_id=838312sasa25&channel_id=C07AsasaASXMH&channel_name=product&timestamp=1438203624.000004&user_id=U03PTHYLB&user_name=sarath&text=testing"
func SlackMessagePostHandler(c *gin.Context) {
	message := newSlackMessage(c)
	if validateRequest(c) == true {
		err := es.Insert(*message)
		if err != nil {
			c.String(http.StatusOK, "falied dure to ===> %s", err.Error())
		}
		c.JSON(http.StatusOK, *message)
	}

}

func validateRequest(c *gin.Context) bool {
	//channelName := c.PostForm("channel_name")
	//token := c.PostForm("token")
	return true
}

func newSlackMessage(c *gin.Context) *types.SlackMessage {
	channelName := c.PostForm("channel_name")
	timestamp := timeStampFromUnix(c.PostForm("timestamp"))
	userName := c.PostForm("user_name")
	text := c.PostForm("text")
	return &types.SlackMessage{Text: text, Channel: channelName, User: userName, Timestamp: &timestamp}
}

func timeStampFromUnix(unixt string) time.Time {
	i, err := strconv.ParseFloat(unixt, 10)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(int64(i), 0)
	return tm
}
