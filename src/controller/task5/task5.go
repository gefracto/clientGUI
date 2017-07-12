package task5

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/clientGUI/src/model"
)

func T5(min, max string) string {
	var data model.T5
	var err error
	var body []byte
	var resp *http.Response
	data.Min, _ = strconv.Atoi(min)
	data.Max, _ = strconv.Atoi(max)

	if body, err = json.Marshal(data); err != nil {
		return "Ошибка маршалинга"
	}

	if resp, err = http.Post("http://localhost:1111/task/5", "application/json", bytes.NewBuffer(body)); err != nil {
		res := "Reason: " + error.Error(err)
		return res
	}
	r, _ := ioutil.ReadAll(resp.Body)
	res := "Result5: \n---------------------\n" + string(r)
	return res
}
