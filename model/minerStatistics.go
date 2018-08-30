package model

import (
	_ "github.com/lib/pq"
)

//在池矿机数量
func GetMinerCountIn(num int) (count int, err error) {

	err = Db.QueryRow("select count(*) from miner where has_join=1 and julianday(datetime('now','localtime'))-julianday(in_time) < ?", num).Scan(&count)
	return
}

//新增矿机数量
func GetMinerCountNew(num int) (count int, err error) {

	err = Db.QueryRow("select count(*) from miner where julianday(datetime('now','localtime'))-julianday(in_time) < ?", num).Scan(&count)
	return
}

//退出矿机数量
func GetMinerCountExit(num int) (count int, err error) {

	err = Db.QueryRow("select count(*) from miner where exit_status=2 and julianday(datetime('now','localtime'))-julianday(in_time) < ?", num).Scan(&count)
	return
}
