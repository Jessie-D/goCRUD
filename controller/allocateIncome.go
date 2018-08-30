package controller

import (
	"net/http"
	"uic-pools/model"

	_ "github.com/lib/pq"
)

// Create a allocateIncome
// POST /allocateIncome/
func NewAllocateIncome(w http.ResponseWriter, r *http.Request) {
	var allocateIncome model.AllocateIncome
	setReqParam(r, &allocateIncome)

	err := allocateIncome.Create()
	if err != nil {
		errorMessage(w, r, err, "创建失败")
		return
	}
	err = Res(w, r, "1", "success", nil)

	return
}
