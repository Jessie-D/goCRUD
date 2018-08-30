package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/jeanphorn/log4go"
)

type res struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func setReqParam(r *http.Request, reqObj interface{}) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, &reqObj)
	reqParam, err := json.MarshalIndent(&reqObj, "", "\t\t")

	log.Info("req:\n %v %v %v \n %v", r.URL, string(r.Method), r.Header.Get("Content-Type"), string(reqParam))
	return
}

func Res(w http.ResponseWriter, r *http.Request, code string, msg string, data interface{}) (err error) {

	dd := res{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	output, err := json.MarshalIndent(&dd, "", "\t\t")
	if err != nil {
		log.Info("生成json出错")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)

	log.Info("res:\n %v %v %v \n %v", r.URL, string(r.Method), r.Header.Get("Content-Type"), string(output))
	return
}

// Convenience function for printing to stdout
func P(a ...interface{}) {
	fmt.Println(a)
}

func errorMessage(w http.ResponseWriter, r *http.Request, err error, msg string) (errback error) {
	var dd res
	if len(msg) > 0 {
		dd = res{
			Code: "500",
			Msg:  msg,
			Data: nil,
		}
	} else {
		dd = res{
			Code: "500",
			Msg:  err.Error(),
			Data: nil,
		}
	}
	output, errback := json.MarshalIndent(&dd, "", "\t\t")

	if errback != nil {
		log.LOGGER("ERROR").Error("生成json出错")
		return
	}

	//w.WriteHeader(500)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	log.LOGGER("ERROR").Error("res:\n %v %v %v \n %v", r.URL, string(r.Method), r.Header.Get("Content-Type"), string(output))

	return
}

func respError(w http.ResponseWriter, error string, code int) {
	dd := res{
		Code: "500",
		Msg:  error,
		Data: nil,
	}
	output, _ := json.MarshalIndent(&dd, "", "\t\t")
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

// version
func version() string {
	return "0.1"
}
