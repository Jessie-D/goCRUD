package controller

import (
	"net/http"
	"path"
	"strconv"
	"uic-pools/model"

	log "github.com/jeanphorn/log4go"
)

// main handlePoolr function
func HandlePoolRequest(w http.ResponseWriter, r *http.Request) {

	var err error
	switch r.Method {
	case "GET":
		err = handlePoolGet(w, r) //查询矿池详情
	case "POST":
		err = handlePoolPost(w, r) //新增矿池
	case "PUT":
		err = handlePoolPut(w, r) //修改费率
	case "DELETE":
		err = handlePoolDelete(w, r) //删除矿池
	}
	if err != nil {
		//errorMessage(w, r, err, err.Error())
		log.LOGGER("ERROR").Error(r.URL, err.Error())
		//respError(w, err.Error(), 500)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Retrieve pools
// GET /pool
func HandlePoolGetAll(w http.ResponseWriter, r *http.Request) {

	pools, err := model.Pools()
	if err != nil {
		log.LOGGER("ERROR").Error(r.URL, err.Error())
		respError(w, err.Error(), 500)
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = Res(w, r, "1", "success", pools)

	return
}

// Retrieve a pool
// GET /pool/1
func handlePoolGet(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		//err = errorMessage(w, r, err, "没有查询id")
		return
	}
	pool, err := model.Retrieve(id)
	if err != nil {
		err = errorMessage(w, r, err, "没有这条记录")
		return
	}
	err = Res(w, r, "1", "success", pool)

	return
}

// Create a pool
// POST /pool/
func handlePoolPost(w http.ResponseWriter, r *http.Request) (err error) {
	// 方式一："contentype application/json"
	var pool model.Pool
	setReqParam(r, &pool)
	if isRate := rateValidate(pool.Rate); isRate != true {
		errorMessage(w, r, err, "费率范围为0-100")
		return
	}

	id, err := pool.Create()
	if err != nil {
		errorMessage(w, r, err, "创建失败")
		return
	}
	err = Res(w, r, "1", "success", id)

	return
}
func rateValidate(rate int) (valid bool) {
	valid = true
	if rate > 100 {
		valid = false
	}
	if rate < 0 {
		valid = false
	}
	return
}

// Update a pool
// PUT /pool/1
func handlePoolPut(w http.ResponseWriter, r *http.Request) (err error) {
	var poolParam model.Pool
	setReqParam(r, &poolParam)

	if isRate := rateValidate(poolParam.Rate); isRate != true {
		errorMessage(w, r, err, "费率范围为0-100")
		return
	}
	pool, err := model.Retrieve(poolParam.Id)
	pool.Rate = poolParam.Rate
	if err != nil {
		err = errorMessage(w, r, err, "没有这条记录")
		return
	}
	err = pool.Update()
	if err != nil {
		errorMessage(w, r, err, "更新失败")
		return
	}
	err = Res(w, r, "1", "success", nil)

	return
}

// Delete a pool
// DELETE /pool/1
func handlePoolDelete(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))

	if err != nil {
		err = Res(w, r, "0", "请传入id,id为数值类型", nil)
		return
	}
	pool, err := model.Retrieve(id)
	if err != nil {
		errorMessage(w, r, err, "没有这条记录")
		return
	}
	err = pool.Delete()
	if err != nil {
		return
	}
	err = Res(w, r, "1", "success", nil)
	if err != nil {
		errorMessage(w, r, err, "服务端异常")
		return
	}
	return
}
