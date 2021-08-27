package model

import "time"

type Shop struct {
	Id int `json:"id"`
	UserId int `json:"user_id"`
	Name string `json:"name"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

func (Shop)TableName()string  {
	return "shop"
}