package server

import (
	"TempBackend/backend"
	"TempBackend/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func (s *Server) PostUserDaily(c *gin.Context) {
	println(time.Now().Format("2006-01-02 15:04:05") + " | " + c.Request.Host + " | " + "post user daily")
	var user model.UserDaily
	err := c.BindJSON(&user)
	if err != nil {
		errMsg := "modify UserDaily failed"
		println(errMsg)
		c.JSON(http.StatusBadRequest, CreateFailureJsonResp(errMsg))
		return
	}
	backend.PostUserDaily(&user)
}
func (s *Server) GetUserStats(c *gin.Context) {
	println(time.Now().Format("2006-01-02 15:04:05") + " | " + c.Request.Host + " | " + "get user status")
	name := strings.TrimSpace(c.Query("name"))
	if name == "" {
		errMsg := "input name is empty"
		println(errMsg)
		c.JSON(http.StatusBadRequest, CreateFailureJsonResp(errMsg))
		return
	}
	userStats := backend.GetUserStats(name)
	if userStats == nil {
		c.JSON(http.StatusInternalServerError, CreateFailureJsonResp("internal failure"))
		return
	}
	c.JSON(http.StatusOK, CreateSuccessJsonResp(userStats))
}

func (s *Server) GetManageDaily(c *gin.Context) {
	println(time.Now().Format("2006-01-02 15:04:05") + " | " + c.Request.Host + " | " + "get manage daily")
	date := strings.TrimSpace(c.Query("date"))
	if date == "" {
		errMsg := "input date is empty"
		println(errMsg)
		c.JSON(http.StatusBadRequest, CreateFailureJsonResp(errMsg))
		return
	}
	manageDaily := backend.GetManageDaily(date)
	if manageDaily == nil {
		c.JSON(http.StatusInternalServerError, CreateFailureJsonResp("internal failure"))
		return
	}
	c.JSON(http.StatusOK, CreateSuccessJsonResp(manageDaily))
}

func (s *Server) GetManageMoon(c *gin.Context) {
	println(time.Now().Format("2006-01-02 15:04:05") + " | " + c.Request.Host + " | " + "get manage all")
	date := strings.TrimSpace(c.Query("date"))
	if len(strings.Split(date, "-")) != 2 {
		errMsg := "wrong input date format, input: " + date
		println(errMsg)
		c.JSON(http.StatusBadRequest, CreateFailureJsonResp(errMsg))
		return
	}
	manageAll := backend.GetManageMoon(date)
	if manageAll == nil {
		c.JSON(http.StatusInternalServerError, CreateFailureJsonResp("internal failure"))
		return
	}
	c.JSON(http.StatusOK, CreateSuccessJsonResp(manageAll))
}

func (s *Server) GetManageAll(c *gin.Context) {
	println(time.Now().Format("2006-01-02 15:04:05") + " | " + c.Request.Host + " | " + "get manage all")
	manageAll := backend.GetManageAll()
	if manageAll == nil {
		c.JSON(http.StatusInternalServerError, CreateFailureJsonResp("internal failure"))
		return
	}
	c.JSON(http.StatusOK, CreateSuccessJsonResp(manageAll))
}

type SuccessJsonResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type FailureJsonResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func CreateSuccessJsonResp(data interface{}) SuccessJsonResp {
	return SuccessJsonResp{
		Code: http.StatusOK,
		Msg:  "success",
		Data: data,
	}
}

func CreateFailureJsonResp(msg string) FailureJsonResp {
	return FailureJsonResp{
		Code: http.StatusInternalServerError,
		Msg:  msg,
	}
}
