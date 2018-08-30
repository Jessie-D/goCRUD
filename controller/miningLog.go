package controller

import (
	"net/http"
	"uic-pools/model"

	log "github.com/jeanphorn/log4go"
	_ "github.com/lib/pq"
)

// Retrieve miningLog
// GET /mininingLogList
func MininingLogList(w http.ResponseWriter, r *http.Request) {

	miningLogs, err := model.MiningLogs()
	if err != nil {
		log.LOGGER("ERROR").Error(r.URL, err.Error())
		respError(w, err.Error(), 500)
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = Res(w, r, "1", "success", miningLogs)

	return
}
func StartMining(w http.ResponseWriter, r *http.Request) {
	var miningLog model.MiningLog
	setReqParam(r, &miningLog)

	id, err := miningLog.Create()
	if err != nil {
		errorMessage(w, r, err, "更新失败")
		return
	}
	err = Res(w, r, "1", "success", id)

	return

}

// Update a miningLog
// Post /StopMining
func StopMining(w http.ResponseWriter, r *http.Request) {

	var miningLog model.MiningLog
	setReqParam(r, &miningLog)

	miningLog, err := model.GetMingLogByID(miningLog.Id)
	if err != nil {
		err = errorMessage(w, r, err, "没有这条记录")
		return
	}
	err = miningLog.Update()
	if err != nil {
		errorMessage(w, r, err, "更新失败")
		return
	}
	err = Res(w, r, "1", "success", nil)

	return
}
