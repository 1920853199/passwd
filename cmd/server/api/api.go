package api

import (
	"encoding/json"
	"net/http"

	"github.com/1920853199/passwd/service"
	"github.com/1920853199/passwd/util"
)

func Execute(w http.ResponseWriter, r *http.Request) {

	resp := service.NewResponse()

	err := util.CheckJwt(r)
	if err != nil {
		resp.Code = 500
		resp.Msg = err.Error()
		util.JSON(w, http.StatusOK, resp)
		return
	}

	params := service.ExecuteParams{}
	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		resp.Code = 500
		resp.Msg = err.Error()
		util.JSON(w, http.StatusOK, resp)
		return
	}
	data, err := service.Parse(params)
	if err != nil {
		resp.Code = 500
		resp.Msg = err.Error()
		util.JSON(w, http.StatusOK, resp)
		return
	}
	resp.Data = data
	util.JSON(w, http.StatusOK, resp)
}
