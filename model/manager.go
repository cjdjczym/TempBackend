package model

// ManageDaily 根据日期查询当天的用户体温信息
type ManageDaily struct {
	NormalCount   string   `json:"normal_count"`
	NormalNames   []string `json:"normal_names"`
	AbnormalCount string   `json:"abnormal_count"`
	AbnormalNames []string `json:"abnormal_names"`
}

// ManageAll 查询所有的体温异常信息
type ManageAll struct {
	NormalCount   string   `json:"normal_count"`
	Normals       []Single `json:"normals"`
	AbnormalCount string   `json:"abnormal_count"`
	Abnormals     []Single `json:"abnormals"`
}

type Single struct {
	Name string `json:"name"`
	Date string `json:"date"`
}
