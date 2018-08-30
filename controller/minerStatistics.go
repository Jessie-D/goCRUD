package controller

import (
	"net/http"
	"path"
	"strconv"
	"uic-pools/model"

	log "github.com/jeanphorn/log4go"
	_ "github.com/lib/pq"
)

type MinerStatisticsParam struct {
	Days           int `json:"days"`
	InCount        int `json:"inCount"`
	OutCount       int `json:"outCount"`
	NewlyIncreased int `json:"newlyIncreased"`
}

func MinerStatistics(w http.ResponseWriter, r *http.Request) {

	var resData MinerStatisticsParam
	num, err := strconv.Atoi(path.Base(r.URL.Path))
	resData.Days = num
	resData.InCount, err = model.GetMinerCountIn(num)
	resData.NewlyIncreased, err = model.GetMinerCountNew(num)
	resData.OutCount, err = model.GetMinerCountExit(num)
	if err != nil {
		log.LOGGER("ERROR").Error(r.URL, err.Error())
		respError(w, err.Error(), 500)
		return
	}
	err = Res(w, r, "1", "success", resData)

	return
}
