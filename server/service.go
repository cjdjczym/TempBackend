package server

import (
	"../backend"
	"../model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (s *Server) PostUserDaily(c *gin.Context) {
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
	name := strings.TrimSpace(c.Param("name"))
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
	date := strings.TrimSpace(c.Param("date"))
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

func (s *Server) GetManageAll(c *gin.Context) {
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
