package task1

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/clientGUI/src/model"
)

func T1(width, height, symbol string) string {
	var data model.T1
	var err error
	var body []byte
	var resp *http.Response
	if data.Width, err = strconv.Atoi(width); err != nil {
		return "Ошибка! Значения должны быть типа int"
	}

	if data.Height, err = strconv.Atoi(height); err != nil {
		return "Ошибка! Значения должны быть типа int"
	}

	data.Symbol = symbol
	if body, err = json.Marshal(data); err != nil {
		return "Ошибка маршалинга"
	}

	if resp, err = http.Post("http://localhost:1111/task/1", "application/json", bytes.NewBuffer(body)); err != nil {
		res := "Reason: " + error.Error(err)
		return res
	}
	r, _ := ioutil.ReadAll(resp.Body)
	res := "Result1: \n---------------------\n" + string(r)
	return res

}
