package model

// UserDaily 用户每日上传数据
type UserDaily struct {
	Id     int    `gorm:"primary_key" json:"id"`
	Name   string `json:"name"`
	Date   string `json:"date"`   // yyyy-mm-dd
	Normal bool   `json:"normal"` // 今日体温是否正常
	Status Status `json:"status"`
}

type Status struct {
	Min    string `json:"min"`    // 最低温度
	Max    string `json:"max"`    // 最高温度
	Avg    string `json:"avg"`    // 平均温度
	Center string `json:"center"` // 中心温度
}

// UserStats 用户总共的正常/异常状态信息
type UserStats struct {
	NoAbnormal bool        `json:"no_abnormal"` // 该用户是否从来没有异常体温
	DailyStats []DailyStat `json:"daily_stats"`
}

type DailyStat struct {
	Date   string `json:"date"`
	Normal bool   `json:"normal"`
}
