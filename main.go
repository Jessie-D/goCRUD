package main

import (
	"net/http"
	"time"
	"uic-pools/controller"
	"uic-pools/model"

	log "github.com/jeanphorn/log4go"
	_ "github.com/mattn/go-sqlite3"
	"github.com/robfig/cron"
)

var config = model.ConfigData

func main() {

	log.Info("uic-pool started at %s", config.Address)
	c := cron.New()
	c.Start()
	c.AddFunc("@midnight", model.ProcessMinerExit)

	server := http.Server{
		Addr:         config.Address, //":8080",
		ReadTimeout:  time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout: time.Duration(config.WriteTimeout * int64(time.Second)),
	}
	http.HandleFunc("/pool/", controller.HandlePoolRequest)
	http.HandleFunc("/pools", controller.HandlePoolGetAll)

	http.HandleFunc("/miner/", controller.HandleMinerRequest)
	http.HandleFunc("/miners", controller.HandleMinerGetAll)
	http.HandleFunc("/minerApplyExit/", controller.HandleApplyExit)

	http.HandleFunc("/statistics/miners/", controller.MinerStatistics)

	http.HandleFunc("/mininingLogList", controller.MininingLogList)
	http.HandleFunc("/startMining", controller.StartMining)
	http.HandleFunc("/stopMining", controller.StopMining)

	http.HandleFunc("/newPoolMining", controller.NewPoolMining)

	http.HandleFunc("/newAllocateIncome", controller.NewAllocateIncome)

	server.ListenAndServe()

	select {}
}
