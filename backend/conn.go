package backend

import (
	"TempBackend/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func connect(cfg *model.Config) *gorm.DB {
	dsn := cfg.DB.User + ":" + cfg.DB.PassWd + "@tcp(" + cfg.DB.Addr + ")/" + cfg.DB.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		println("connect to mysql failed, err: " + err.Error())
		return nil
	}
	err = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.UserDaily{})
	if err != nil {
		println("db auto migrate failed, err: " + err.Error())
		return nil
	}
	return db
}

func PostUserDaily(user *model.UserDaily, cfg *model.Config) {
	db := connect(cfg)
	err := db.Create(user).Error
	if err != nil {
		println("post user daily failed, err: " + err.Error())
		return
	}
}

func GetUserStats(name string, cfg *model.Config) *model.UserStats {
	var users []model.UserDaily
	db := connect(cfg)
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

func GetManageDaily(date string, cfg *model.Config) *model.ManageDaily {
	var users []model.UserDaily
	db := connect(cfg)
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

func GetManageMoon(date string, cfg *model.Config) *model.ManageAll {
	var users []model.UserDaily
	db := connect(cfg)
	err := db.Where("date like ?", date+"%").Find(&users).Error
	if err != nil {
		println("get manage moon failed, err: " + err.Error())
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

func GetManageAll(cfg *model.Config) *model.ManageAll {
	var users []model.UserDaily
	db := connect(cfg)
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

func GetUserCount(cfg *model.Config) int64 {
	var count int64
	db := connect(cfg)
	err := db.Table("user_dailies").Select("count(distinct(name))").Count(&count).Error
	if err != nil {
		println("get user count failed, err: " + err.Error())
		return 0
	}
	return count
}

func GetAllNormal(cfg *model.Config) int64 {
	var count int64
	db := connect(cfg)
	err := db.Table("user_dailies").Select("count(*)").Where("normal = ?", true).Count(&count).Error
	if err != nil {
		println("get all normal failed, err: " + err.Error())
		return 0
	}
	return count
}

func GetAllAbnormal(cfg *model.Config) int64 {
	var count int64
	db := connect(cfg)
	err := db.Table("user_dailies").Select("count(*)").Where("normal = ?", false).Count(&count).Error
	if err != nil {
		println("get all abnormal failed, err: " + err.Error())
		return 0
	}
	return count
}

func GetTodayNormal(cfg *model.Config) int64 {
	var count int64
	db := connect(cfg)
	err := db.Table("user_dailies").Select("count(*)").Where("normal = ? AND date = ?", true, time.Now().Format("2006-01-02")).Count(&count).Error
	if err != nil {
		println("get today normal failed, err: " + err.Error())
		return 0
	}
	return count
}

func GetTodayAbnormal(cfg *model.Config) int64 {
	var count int64
	db := connect(cfg)
	err := db.Table("user_dailies").Select("count(*)").Where("normal = ? AND date = ?", false, time.Now().Format("2006-01-02")).Count(&count).Error
	if err != nil {
		println("get today abnormal failed, err: " + err.Error())
		return 0
	}
	return count
}
