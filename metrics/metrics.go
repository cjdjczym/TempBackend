package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	AvgTempGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "avg_temp_gauge",
			Help: "avg temp gauge",
		})
	UsersGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "users_gauge",
			Help: "users gauge",
		})
	NormalGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "normal_total_gauge",
			Help: "normal total gauge",
		})
	AbnormalGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "abnormal_total_gauge",
			Help: "abnormal total gauge",
		})
	NormalDailyGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "normal_daily_gauge",
			Help: "normal daily gauge",
		})
	AbnormalDailyGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "abnormal_daily_gauge",
			Help: "abnormal daily gauge",
		})
	PostUserDailyCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "post_user_daily_counter",
			Help: "post user daily counter",
		})
	GetUserStatsCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "get_user_stats_counter",
			Help: "get user stats counter",
		})
	GetManageDailyCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "get_manage_daily_counter",
			Help: "get manage daily counter",
		})
	GetManageMoonCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "get_manage_moon_counter",
			Help: "get manage moon counter",
		})
	GetManageAllCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "get_manage_all_counter",
			Help: "get manage all counter",
		})
)

func RegisterProxyMetrics() {
	prometheus.MustRegister(AvgTempGauge)
	prometheus.MustRegister(UsersGauge)
	prometheus.MustRegister(NormalGauge)
	prometheus.MustRegister(AbnormalGauge)
	prometheus.MustRegister(NormalDailyGauge)
	prometheus.MustRegister(AbnormalDailyGauge)

	prometheus.MustRegister(PostUserDailyCounter)
	prometheus.MustRegister(GetUserStatsCounter)
	prometheus.MustRegister(GetManageDailyCounter)
	prometheus.MustRegister(GetManageMoonCounter)
	prometheus.MustRegister(GetManageAllCounter)
}
