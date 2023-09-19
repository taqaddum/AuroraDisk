package model

import (
	"log"
	"xorm.io/xorm"
)
import _ "github.com/lib/pq"

var DB *xorm.Engine

func init() {
	var dbUrl string
	DB, _ = xorm.NewEngine("postgres", dbUrl)
	if err := DB.Ping(); err != nil {
		log.Fatalln("数据库连接失败", err.Error())
	}
}

func DisConnect() {
	if err := DB.Close(); err != nil {
		log.Println("数据库关闭失败orz...", err.Error())
	}
}
