package model

import (
	"database/sql"
	"time"

	log "github.com/jeanphorn/log4go"
	_ "github.com/lib/pq"
)

type PoolMining struct {
	Id          int       `json:"id"`
	UDID        string    `json:"UDID"`
	PoolID      int       `json:"poolID"`
	BlockNumber int       `json:"blockNumber"`
	Date        time.Time `json:"date"`
	Income      float64   `json:"income"`
}

func clearTransaction(tx *sql.Tx) {
	err := tx.Rollback()
	if err != sql.ErrTxDone && err != nil {
		log.LOGGER("ERROR").Error("rollback error %v", err.Error())
	}
}

func AllocateIncomes(miners []Miner, miningId int) (err error) {
	tx, err := Db.Begin()
	if err != nil {
		log.LOGGER("ERROR").Error("事务begin error %v ", err.Error())
		return
	}
	defer clearTransaction(tx)
	for _, miner := range miners {
		var allocateIncomeObj = AllocateIncome{
			MiningId: miningId,
			UDID:     miner.UDID,
			PoolID:   miner.PoolID,
			Income:   1,
			Fee:      1,
			Date:     time.Now(),
			Status:   1,
		}
		err = allocateIncomeObj.Create()
		if err != nil {
			log.LOGGER("ERROR").Error("分配出错%v", err.Error())
		}
	}

	if err := tx.Commit(); err != nil {
		log.LOGGER("ERROR").Error("commit error %v", err.Error())

	}
	log.Info("分配事务成功")
	return
}

// Create a new poolMining
func (poolMining *PoolMining) Create() (id int64, err error) {
	//插入数据
	stmt, err := Db.Prepare("INSERT INTO poolMining(UDID, pool_id,block_number,income,date)  values(?,?,?,?,?)")
	if err != nil {
		return
	}

	res, err := stmt.Exec(poolMining.UDID, poolMining.PoolID, poolMining.BlockNumber, poolMining.Income, time.Now().Unix())
	if err != nil {
		return
	}

	id, err = res.LastInsertId()

	if err != nil {
		return
	}

	return
}
