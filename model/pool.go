package model

import (
	"time"

	_ "github.com/lib/pq"
)

type Pool struct {
	Id         int       `json:"id"`
	UDID       string    `json:"UDID"`
	Name       string    `json:"name"`
	Rate       int       `json:"rate"`       //管理费
	FeeAccount string    `json:"feeAccount"` //管理费账户
	Account    string    `json:"account"`    //收益账户
	Password   string    `json:"password"`
	Grade      string    `json:"grade"`
	CreatedAt  time.Time `json:"createdAt"`
}

func Pools() (pools []Pool, err error) {
	//查询数据
	rows, err := Db.Query("SELECT * FROM pools")
	if err != nil {
		return
	}

	for rows.Next() {
		conv := Pool{}
		if err = rows.Scan(&conv.Id, &conv.UDID, &conv.Name, &conv.Rate, &conv.FeeAccount, &conv.Account, &conv.Password, &conv.Grade, &conv.CreatedAt); err != nil {
			return
		}
		pools = append(pools, conv)
	}
	rows.Close()
	return
}

// Get a single pool
func Retrieve(id int) (conv Pool, err error) {
	conv = Pool{}
	err = Db.QueryRow("select * from pools where id = $1", id).Scan(&conv.Id, &conv.UDID, &conv.Name, &conv.Rate, &conv.FeeAccount, &conv.Account, &conv.Password, &conv.Grade, &conv.CreatedAt)
	return
}
func GetPoolByUDID(udid string) (conv Pool, err error) {
	conv = Pool{}
	err = Db.QueryRow("select * from pools where UDID = $1", udid).Scan(&conv.Id, &conv.UDID, &conv.Name, &conv.Rate, &conv.FeeAccount, &conv.Account, &conv.Password, &conv.Grade, &conv.CreatedAt)
	return
}

// Create a new pool
func (pool *Pool) Create() (id int64, err error) {
	//插入数据
	stmt, err := Db.Prepare("INSERT INTO pools(UDID, name, rate,feeAccount ,account,password,created_at,grade)  values(?,?,?,?,?,?,?,1)")
	if err != nil {
		return
	}

	res, err := stmt.Exec("", pool.Name, pool.Rate, pool.FeeAccount, pool.Account, pool.Password, time.Now().Format(dateTimeLayout))
	if err != nil {
		return
	}

	id, err = res.LastInsertId()

	if err != nil {
		return
	}

	return
}

// Update a pool
func (pool *Pool) Update() (err error) {
	stmt, err := Db.Prepare("update pools set rate=? , password=? ,feeAccount =?  where id=?")
	_, err = stmt.Exec(pool.Rate, pool.Password, pool.FeeAccount, pool.Id)

	return
}

// Delete a pool
func (pool *Pool) Delete() (err error) {
	_, err = Db.Exec("delete from pools where id = $1", pool.Id)
	return
}
