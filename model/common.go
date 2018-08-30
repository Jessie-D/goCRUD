package model

import (
	"encoding/json"
	"os"

	"database/sql"

	log "github.com/jeanphorn/log4go"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB
var i = 0

const (
	dateTimeLayout = "2006-01-02 15:04:05"
	dateLayout     = "2006-01-02"
)

type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
}

var ConfigData Configuration

func init() {
	os.MkdirAll("log/info", 0777)
	os.MkdirAll("log/error", 0777)
	// 加载日志配置文件
	log.LoadConfiguration("./config/log4go.json")
	file, err := os.Open("./config/sever.json")
	if err != nil {
		log.LOGGER("ERROR").Error("Cannot open config file", err)
	}
	decoder := json.NewDecoder(file)
	ConfigData = Configuration{}
	err = decoder.Decode(&ConfigData)
	if err != nil {
		log.LOGGER("ERROR").Error("Cannot get configuration from file", err)
	}

	Db, err = sql.Open("sqlite3", "pool.db")
	if err != nil {
		log.LOGGER("ERROR").Error("连接数据库失败", err.Error())

	}
	log.Info("成功连接数据库")

}
