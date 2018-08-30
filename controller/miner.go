package controller

import (
	"net/http"
	"path"
	"strconv"
	"uic-pools/model"

	log "github.com/jeanphorn/log4go"
	_ "github.com/lib/pq"
)

type AddParam struct {
	model.Miner
	Password string `json:"password"`
}

// main handleMinerr function
func HandleMinerRequest(w http.ResponseWriter, r *http.Request) {

	var err error
	switch r.Method {
	case "GET":
		err = handleMinerGet(w, r) //查询矿机详情
	case "POST":
		err = handleMinerPost(w, r) //新增矿机
	case "PUT":
		err = handleMinerPut(w, r) //修改费率
	case "DELETE":
		err = handleMinerDelete(w, r) //删除矿机
	}
	if err != nil {
		//errorMessage(w, r, err, err.Error())
		log.LOGGER("ERROR").Error(r.URL, err.Error())
		//respError(w, err.Error(), 500)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Retrieve miners
// GET /pool
func HandleMinerGetAll(w http.ResponseWriter, r *http.Request) {

	miner, err := model.Miners()
	if err != nil {
		log.LOGGER("ERROR").Error(r.URL, err.Error())
		respError(w, err.Error(), 500)
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = Res(w, r, "1", "success", miner)

	return
}

// get
// minerApplyExit/:id
func HandleApplyExit(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	err = model.ApplyExit(id)
	if err != nil {
		log.LOGGER("ERROR").Error(r.URL, err.Error())
		respError(w, err.Error(), 500)
		return
	}
	err = Res(w, r, "1", "success", nil)

	return
}

// Retrieve a miner
// GET /miner/1
func handleMinerGet(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		err = errorMessage(w, r, err, "没有查询id")
		return
	}
	miner, err := model.GetMinerById(id)
	if err != nil {
		err = errorMessage(w, r, err, "没有这条记录")
		return
	}
	err = Res(w, r, "1", "success", miner)

	return
}

// Create a miner
// POST /miner/
func handleMinerPost(w http.ResponseWriter, r *http.Request) (err error) {

	var miner AddParam
	setReqParam(r, &miner)

	//是否已加入某个矿池验证
	var count int64

	err = model.Db.QueryRow("SELECT count(*)  FROM miner  where UDID=? and has_join=1", miner.UDID).Scan(&count)
	if err != nil {
		errorMessage(w, r, err, err.Error())
		return
	}
	if count == 1 {
		errorMessage(w, r, err, "此矿机已加入一个矿池，不可再加入")
		return
	}
	//矿池是否设置加密开关

	pool, err := model.Retrieve(miner.PoolID)
	//根据pool_id查询矿池信息
	if err != nil {
		errorMessage(w, r, err, err.Error())
		return
	}
	if pool.Password != miner.Password {
		errorMessage(w, r, err, "矿池密码输入错误")
		return
	}

	//矿池矿机数量限制的校验
	err = model.Db.QueryRow("SELECT count(*)  FROM miner").Scan(&count)
	if err != nil {
		return
	}
	if count == 100 {
		errorMessage(w, r, err, "此矿池已满100台矿机，不可再加入")
		return
	}

	id, err := miner.Create()
	if err != nil {
		errorMessage(w, r, err, "创建失败")
		return
	}
	err = Res(w, r, "1", "success", id)

	return
}

// Update a miner
// PUT /miner/1
func handleMinerPut(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		err = errorMessage(w, r, err, "没有查询id")
		return
	}
	miner, err := model.GetMinerById(id)
	if err != nil {
		err = errorMessage(w, r, err, "没有这条记录")
		return
	}

	setReqParam(r, &miner)

	err = miner.Update()

	if err != nil {
		errorMessage(w, r, err, "更新失败")
		return
	}
	err = Res(w, r, "1", "success", nil)

	return
}

// Delete a miner
// DELETE /miner/1
func handleMinerDelete(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))

	if err != nil {
		err = Res(w, r, "0", "请传入id id为数值类型", nil)
		return
	}
	miner, err := model.GetMinerById(id)
	if err != nil {
		err = errorMessage(w, r, err, "没有这条记录")

		return
	}
	err = miner.Delete()
	if err != nil {
		return
	}
	err = Res(w, r, "1", "success", nil)
	if err != nil {
		err = errorMessage(w, r, err, "服务端异常")
		return
	}
	return
}
