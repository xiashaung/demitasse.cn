package model

import "time"

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Head string `json:"head"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

func (User)TableName()string  {
	return "user"
}
