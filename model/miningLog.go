package model

import (
	"time"

	log "github.com/jeanphorn/log4go"
	_ "github.com/lib/pq"
)

type MiningLog struct {
	Id         int       `json:"id"`
	UDID       string    `json:"UDID"`
	PoolID     int       `json:"poolID"`
	StartTime  time.Time `json:"startTime"`
	StopTime   time.Time `json:"stopTime"`
	MiningTime int       `json:"miningTime"`
}

// // Get a single miningLog
func GetMingLogByID(id int) (conv MiningLog, err error) {
	conv = MiningLog{}
	err = Db.QueryRow("select * from miningLog where id = $1", id).Scan(&conv.Id, &conv.UDID, &conv.PoolID, &conv.StartTime, &conv.StopTime, &conv.MiningTime)
	return
}

// Create a new miningLog
func (miningLog *MiningLog) Create() (id int64, err error) {
	//插入数据
	stmt, err := Db.Prepare("INSERT INTO miningLog(UDID, pool_id, start_time,stop_time,mining_time)  values(?,?,?,?,?)")
	if err != nil {
		return
	}

	res, err := stmt.Exec(miningLog.UDID, miningLog.PoolID, time.Now().Unix(), "", 0)
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

// Update a miningLog
func (miningLog *MiningLog) Update() (err error) {

	var stopTime = time.Now().Unix()
	var miningTime = stopTime - miningLog.StartTime.Unix()
	stmt, err := Db.Prepare("update miningLog set stop_time=?,mining_time=? where id=?")
	_, err = stmt.Exec(stopTime, miningTime, miningLog.Id)

	return
}

// Delete a miningLog
func (miningLog *MiningLog) Delete() (err error) {
	_, err = Db.Exec("delete from miningLog where id = $1", miningLog.Id)
	return
}
func MiningLogs() (miningLog []MiningLog, err error) {
	// 	//查询数据
	rows, err := Db.Query("SELECT * FROM miningLog")
	if err != nil {
		return
	}

	for rows.Next() {
		conv := MiningLog{}
		if err = rows.Scan(&conv.Id, &conv.UDID, &conv.PoolID, &conv.StartTime, &conv.StopTime, &conv.MiningTime); err != nil {
			return
		}
		miningLog = append(miningLog, conv)
	}
	rows.Close()
	return
}
