package service

import (
	"bytes"
	json2 "encoding/json"
	"golang_homework/ch5-7/go_customs/model"
	"io/ioutil"
	"net/http"
)

type GosuslugiService struct {
	Server *http.Server
	Port   string
}

func (service *GosuslugiService) GetTicket(data *model.CheckinServiceRequestModel) string {
	json, err := json2.Marshal(data)
	if err != nil {
		panic(err)
	}
	resp, err := http.Post("http://localhost:"+service.Port, "application/json", bytes.NewBuffer(json))
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	str := string(body)
	return str
}
