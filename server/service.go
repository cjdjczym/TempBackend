package server

import (
	"TempBackend/backend"
	"TempBackend/metrics"
	"TempBackend/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

// @Summary 用户上传体温
// @Produce json
// @Success 200 {string} json "{"code": 200,"msg": "success","data": null}"
// @Router /api/user/daily [post]
func (s *Server) PostUserDaily(c *gin.Context) {
	metrics.PostUserDailyCounter.Inc()

	println(time.Now().Format("2006-01-02 15:04:05") + " | " + c.Request.Host + " | " + "post user daily")
	var user model.UserDaily
	err := c.BindJSON(&user)
	if err != nil {
		errMsg := "modify UserDaily failed"
		println(errMsg)
		c.JSON(http.StatusBadRequest, CreateFailureJsonResp(errMsg))
		return
	}

	backend.PostUserDaily(&user, s.cfg)
	c.JSON(http.StatusOK, CreateSuccessJsonResp(nil))

	var userCount = float64(backend.GetUserCount(s.cfg))
	metrics.UsersGauge.Set(userCount)
	print("user count: ")
	println(userCount)
	var allNormal = float64(backend.GetAllNormal(s.cfg))
	metrics.NormalGauge.Set(allNormal)
	print("all normal: ")
	println(allNormal)
	var allAbnormal = float64(backend.GetAllAbnormal(s.cfg))
	metrics.AbnormalGauge.Set(allAbnormal)
	print("all abnormal: ")
	println(allAbnormal)
	var dailyNormal = float64(backend.GetTodayNormal(s.cfg))
	metrics.NormalDailyGauge.Set(dailyNormal)
	print("today normal: ")
	println(dailyNormal)
	var dailyAbnormal = float64(backend.GetTodayAbnormal(s.cfg))
	metrics.AbnormalDailyGauge.Set(dailyAbnormal)
	print("today abnormal: ")
	println(dailyAbnormal)
}

// @Summary 用户获取自己近期体温情况
// @Produce json
// @Success 200 {string} json "{"code": 200,"msg": "success","data": "no_abnormal: true,"daily_stats": ["date": "2021-08-30","normal": true]"}"
// @Router /api/user/stats [get]
func (s *Server) GetUserStats(c *gin.Context) {
	metrics.GetUserStatsCounter.Inc()
	println(time.Now().Format("2006-01-02 15:04:05") + " | " + c.Request.Host + " | " + "get user status")
	name := strings.TrimSpace(c.Query("name"))
	if name == "" {
		errMsg := "input name is empty"
		println(errMsg)
		c.JSON(http.StatusBadRequest, CreateFailureJsonResp(errMsg))
		return
	}
	userStats := backend.GetUserStats(name, s.cfg)
	if userStats == nil {
		c.JSON(http.StatusInternalServerError, CreateFailureJsonResp("internal failure"))
		return
	}
	c.JSON(http.StatusOK, CreateSuccessJsonResp(userStats))
}

// @Summary 根据日期查询当天的用户体温信息
// @Produce json
// @Success 200 {string} json "{"code": 200,"msg": "success","data": "normal_count: "2","normal_names": ["xiaoming","xiaohong"],"abnormal_count": "1","abnormal_names": ["xiaoli"]}"
// @Router /api/manager/daily/{date} [get]
func (s *Server) GetManageDaily(c *gin.Context) {
	metrics.GetManageDailyCounter.Inc()
	println(time.Now().Format("2006-01-02 15:04:05") + " | " + c.Request.Host + " | " + "get manage daily")
	date := strings.TrimSpace(c.Query("date"))
	if date == "" {
		errMsg := "input date is empty"
		println(errMsg)
		c.JSON(http.StatusBadRequest, CreateFailureJsonResp(errMsg))
		return
	}
	manageDaily := backend.GetManageDaily(date, s.cfg)
	if manageDaily == nil {
		c.JSON(http.StatusInternalServerError, CreateFailureJsonResp("internal failure"))
		return
	}
	c.JSON(http.StatusOK, CreateSuccessJsonResp(manageDaily))
}

// @Summary 根据年月查询当月的用户体温信息
// @Produce json
// @Success 200 {string} json "{"code": 200,"msg": "success","data": "normal_count: "2","normals": [{"name": "xiaoming","date": "2021-08-19"},{"name": "xiaohong","date": "2021-08-25"}],"abnormal_count": "1","abnormals": [{"name": "xiaoli","date": "2021-08-12"}]}"
// @Router /api/manager/moon/{date} [get]
func (s *Server) GetManageMoon(c *gin.Context) {
	metrics.GetManageMoonCounter.Inc()
	println(time.Now().Format("2006-01-02 15:04:05") + " | " + c.Request.Host + " | " + "get manage all")
	date := strings.TrimSpace(c.Query("date"))
	if len(strings.Split(date, "-")) != 2 {
		errMsg := "wrong input date format, input: " + date
		println(errMsg)
		c.JSON(http.StatusBadRequest, CreateFailureJsonResp(errMsg))
		return
	}
	manageAll := backend.GetManageMoon(date, s.cfg)
	if manageAll == nil {
		c.JSON(http.StatusInternalServerError, CreateFailureJsonResp("internal failure"))
		return
	}
	c.JSON(http.StatusOK, CreateSuccessJsonResp(manageAll))
}

// @Summary 查询所有的用户体温信息
// @Produce json
// @Success 200 {string} json "{"code": 200,"msg": "success","data": "normal_count: "2","normals": [{"name": "xiaoming","date": "2021-08-19"},{"name": "xiaohong","date": "2021-08-25"}],"abnormal_count": "1","abnormals": [{"name": "xiaoli","date": "2021-08-12"}]}"
// @Router /api/manager/all [get]
func (s *Server) GetManageAll(c *gin.Context) {
	metrics.GetManageAllCounter.Inc()
	println(time.Now().Format("2006-01-02 15:04:05") + " | " + c.Request.Host + " | " + "get manage all")
	manageAll := backend.GetManageAll(s.cfg)
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
