package controller

import (
	"net/http"
	"uic-pools/model"

	log "github.com/jeanphorn/log4go"
	_ "github.com/lib/pq"
)

// Create a poolMining
// POST /poolMining/
func NewPoolMining(w http.ResponseWriter, r *http.Request) {

	var poolMining model.PoolMining
	setReqParam(r, &poolMining)
	id, err := poolMining.Create()
	log.Info("成功加入一条矿池挖矿记录 id=%v", id)
	miners, err := model.GetMinerMing(poolMining.PoolID)
	err = model.AllocateIncomes(miners, int(id))
	if err != nil {
		errorMessage(w, r, err, "分配失败")
		return
	}
	err = Res(w, r, "1", "success", nil)

	return
}
