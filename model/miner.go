package model

import (
	"time"

	log "github.com/jeanphorn/log4go"
	_ "github.com/lib/pq"
)

type Miner struct {
	Id             int       `json:"id"`
	UDID           string    `json:"UDID"`
	PoolID         int       `json:"poolID"`        //加入矿池id
	RemainingTime  int       `json:"remainingTime"` //挖矿剩余时间
	UpdateTime     time.Time `json:"updateTime"`
	Account        string    `json:"account"`        //收益账户
	Status         bool      `json:"status"`         //挖矿状态
	ComputingPower float64   `json:"computingPower"` //算力
	InTime         time.Time `json:"inTime"`
	OutTime        time.Time `json:"outTime"`
	HasJoin        bool      `json:"hasJoin"`    //是否已加入矿池
	ExitStatus     int       `json:"exitStatus"` //退出矿池状态 0空 1申请 2退出 默认0
}

func Miners() (miners []Miner, err error) {
	//查询数据
	rows, err := Db.Query("SELECT * FROM miner")
	if err != nil {
		return
	}

	for rows.Next() {
		conv := Miner{}
		if err = rows.Scan(&conv.Id, &conv.UDID, &conv.PoolID, &conv.Account, &conv.RemainingTime, &conv.UpdateTime, &conv.Status, &conv.InTime, &conv.OutTime, &conv.ComputingPower, &conv.HasJoin, &conv.ExitStatus); err != nil {
			return
		}
		miners = append(miners, conv)
	}
	rows.Close()
	return
}

// Get a single miner
func GetMinerById(id int) (conv Miner, err error) {
	conv = Miner{}
	err = Db.QueryRow("select * from miner where id = $1", id).Scan(&conv.Id, &conv.UDID, &conv.PoolID, &conv.Account, &conv.RemainingTime, &conv.UpdateTime, &conv.Status, &conv.InTime, &conv.OutTime, &conv.ComputingPower, &conv.HasJoin, &conv.ExitStatus)
	return
}

// Create a new miner
func (miner *Miner) Create() (id int64, err error) {

	//插入数据
	stmt, err := Db.Prepare("INSERT INTO miner(UDID, account, pool_ID, status,remaining_time,computing_power,update_time,in_time,out_time,has_join,exit_status)  values(?,?,?,?,?,?,?,?,?,1,0)")
	if err != nil {
		return
	}

	res, err := stmt.Exec(miner.UDID, miner.Account, miner.PoolID, miner.Status, miner.RemainingTime, miner.ComputingPower, time.Now().Unix(), time.Now().Format(dateLayout), "")
	if err != nil {
		return
	}

	id, err = res.LastInsertId()
	log.Info(id)
	if err != nil {
		return
	}

	return
}

// Update a miner
func (miner *Miner) Update() (err error) {
	//_, err = Db.Exec("update miners set topic = $2 where id = ", miner.Id, miner.Name)
	stmt, err := Db.Prepare("update miner set status=? ,remaining_time=?,update_time=? where id=?")
	_, err = stmt.Exec(miner.Status, miner.RemainingTime, miner.UpdateTime, miner.Id)

	return
}

// Delete a miner
func (miner *Miner) Delete() (err error) {
	_, err = Db.Exec("delete from miner where id = $1", miner.Id)
	return
}
func ApplyExit(id int) (err error) {
	stmt, err := Db.Prepare("update miner set exit_status=1 where id=?")
	_, err = stmt.Exec(id)

	return
}

func ProcessMinerExit() {
	stmt, err := Db.Prepare("update miner set exit_status=2 ,has_join=0,out_time=? where exit_status=1")
	log.LOGGER("ERROR").Error(err.Error())
	_, err = stmt.Exec(time.Now().Format(dateLayout))
	log.LOGGER("ERROR").Error(err.Error())
	return
}

func GetMinerMing(poolID int) (miners []Miner, err error) {
	//查询数据
	rows, err := Db.Query("SELECT * FROM miner where pool_id=? and status=1 and has_join=1", poolID)
	if err != nil {
		return
	}

	for rows.Next() {
		conv := Miner{}
		if err = rows.Scan(&conv.Id, &conv.UDID, &conv.PoolID, &conv.Account, &conv.RemainingTime, &conv.UpdateTime, &conv.Status, &conv.InTime, &conv.OutTime, &conv.ComputingPower, &conv.HasJoin, &conv.ExitStatus); err != nil {
			return
		}
		miners = append(miners, conv)
	}
	rows.Close()
	return
}
