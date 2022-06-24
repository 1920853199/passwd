package util

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/1920853199/passwd/service"
)

func JSON(w http.ResponseWriter, code int, resp service.Result) {
	msg, _ := json.Marshal(resp)
	w.Header().Set("content-type", "text/json")
	w.WriteHeader(code)
	w.Write(msg)
}

func Post(u string, form service.ExecuteParams, token string) ([]byte, error) {

	body, err := json.Marshal(form)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", u, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return d, nil
}
