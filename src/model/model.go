package model

import (
	"fmt"
	"github.com/xiashaung/demitasse.cn/lib"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init()  {
	var dsn = fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		lib.Db.User,
		lib.Db.Password,
		lib.Db.Host,
		lib.Db.Port,
		lib.Db.Name,
	)
	DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

//Table
func Table(value interface{}) *gorm.DB {
	DB.Model(value);
	return DB
}