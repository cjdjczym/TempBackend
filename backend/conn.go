package backend

import (
	"../model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	user   = "cj"
	passwd = "010512"
	dbAddr = "127.0.0.1:3306"
	dbName = "test"
)

func connect() *gorm.DB {
	dsn := user + ":" + passwd + "@tcp(" + dbAddr + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		println("connect to mysql failed, err: " + err.Error())
		return nil
	}
	return db
}

func PostUserDaily(user *model.UserDaily) {
	db := connect()
	err := db.Create(user).Error
	if err != nil {
		println("post user daily failed, err: " + err.Error())
		return
	}
}

func GetUserStats(name string) *model.UserStats {
	var users []model.UserDaily
	db := connect()
	err := db.Where("name = ?", name).Find(&users).Error
	if err != nil {
		println("get user stats failed, err: " + err.Error())
		return nil
	}
	userStats := &model.UserStats{NoAbnormal: true, DailyStats: make([]model.DailyStat, len(users))}
	for _, v := range users {
		if !v.Normal {
			userStats.NoAbnormal = false
		}
		stat := model.DailyStat{Date: v.Date, Normal: v.Normal}
		userStats.DailyStats = append(userStats.DailyStats, stat)
	}
	return userStats
}

func GetManageDaily(date string) *model.ManageDaily {
	var users []model.UserDaily
	db := connect()
	err := db.Where("date = ?", date).Find(&users).Error
	if err != nil {
		println("get manage daily failed, err: " + err.Error())
		return nil
	}
	var normalNames, abnormalNames []string
	for _, v := range users {
		if v.Normal {
			normalNames = append(normalNames, v.Name)
		} else {
			abnormalNames = append(abnormalNames, v.Name)
		}
	}
	return &model.ManageDaily{NormalCount: fmt.Sprintf("%d", len(normalNames)), NormalNames: normalNames,
		AbnormalCount: fmt.Sprintf("%d", len(abnormalNames)), AbnormalNames: abnormalNames}
}

func GetManageAll() *model.ManageAll {
	var users []model.UserDaily
	db := connect()
	err := db.Find(&users).Error
	if err != nil {
		println("get manage all failed, err: " + err.Error())
		return nil
	}
	var normals, abnormals []model.Single
	for _, v := range users {
		if v.Normal {
			normals = append(normals, model.Single{Name: v.Name, Date: v.Date})
		} else {
			abnormals = append(abnormals, model.Single{Name: v.Name, Date: v.Date})
		}
	}
	return &model.ManageAll{NormalCount: fmt.Sprintf("%d", len(normals)), Normals: normals,
		AbnormalCount: fmt.Sprintf("%d", len(abnormals)), Abnormals: abnormals}
}
