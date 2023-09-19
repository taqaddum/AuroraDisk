package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"strings"
)

var SysCfg system
var DBCfg database

func Load(path string) {
	viper.SetConfigName("settings")
	viper.SetConfigType("toml")
	viper.AddConfigPath(path)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("无法加载配置文件", err.Error())
	}

	viper.UnmarshalKey("db", &DBCfg)
	{
		params := viper.GetStringMapString("db.parameters")
		DBCfg.SetURL(params)
	}

}

func (db *database) SetURL(params map[string]string) {
	db.URL = fmt.Sprintf("%s://%s:%s@%s:%d/%s", db.Type, db.User, db.Passwd, db.Host, db.Port, db.Name)
	if len(params) != 0 {
		db.URL += "?"
		var tmp []string
		for k, v := range params {
			tmp = append(tmp, k+"="+v)
		}
		db.URL += strings.Join(tmp, "&&")
	}
}
