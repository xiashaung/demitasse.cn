package model

type TalentInfo struct {
	TalentId   int    `json:"talent_id" form:"id"`
	TalentName string `json:"talent_name" form:"name"`
}

func (TalentInfo) TableName() string {
	return "talent_info"
}

func (t *TalentInfo)GetName() string {
	return t.TalentName
}
