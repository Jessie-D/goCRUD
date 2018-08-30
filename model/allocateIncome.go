package model

import (
	"time"

	log "github.com/jeanphorn/log4go"
	_ "github.com/lib/pq"
)

type AllocateIncome struct {
	Id       int       `json:"id"`
	MiningId int       `json:"miningId"`
	UDID     string    `json:"UDID"`
	PoolID   int       `json:"poolID"`
	Income   float64   `json:"income"`
	Fee      float64   `json:"fee"`
	Date     time.Time `json:"date"`
	Status   int       `json:"status"` //状态：1已结算 2未结算
	//Todo add account
}

// Create a new allocateIncome
func (allocateIncome *AllocateIncome) Create() (err error) {
	//插入数据
	stmt, err := Db.Prepare("INSERT INTO allocateIncome(UDID, pool_id,mining_id,income,fee,date,status)  values(?,?,?,?,?,?,?)")
	if err != nil {
		return
	}
	res, err := stmt.Exec(allocateIncome.UDID, allocateIncome.PoolID, allocateIncome.MiningId, allocateIncome.Income, allocateIncome.Fee, time.Now().Unix(), allocateIncome.Status)
	if err != nil {
		return
	}

	id, err := res.LastInsertId()
	log.Info(id)
	if err != nil {
		return
	}

	return
}
