package task2

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/clientGUI/src/model"
)

func T2(a, b, c, d string) string {
	var data [2]model.T2
	var err error
	var body []byte
	var resp *http.Response
	if data[0].Width, err = strconv.Atoi(a); err != nil {
		return "Ошибка! Значения должны быть типа int"
	}
	if data[0].Height, err = strconv.Atoi(b); err != nil {
		return "Ошибка! Значения должны быть типа int"
	}
	if data[1].Width, err = strconv.Atoi(c); err != nil {
		return "Ошибка! Значения должны быть типа int"
	}
	if data[1].Height, err = strconv.Atoi(d); err != nil {
		return "Ошибка! Значения должны быть типа int"
	}

	if body, err = json.Marshal(data); err != nil {
		return "Ошибка маршалинга"
	}

	if resp, err = http.Post("http://localhost:1111/task/2", "application/json", bytes.NewBuffer(body)); err != nil {
		res := "Reason: " + error.Error(err)
		return res
	}

	r, _ := ioutil.ReadAll(resp.Body)
	res := "Result2: \n---------------------\n" + string(r)
	return res
}
