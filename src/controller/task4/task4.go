package task4

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/clientGUI/src/model"
)

func T4(num string) string {
	var data model.T4
	var err error
	var body []byte
	var resp *http.Response
	data.Number, _ = strconv.Atoi(num)
	if body, err = json.Marshal(data); err != nil {
		return "Ошибка маршалинга"
	}

	if resp, err = http.Post("http://localhost:1111/task/4", "application/json", bytes.NewBuffer(body)); err != nil {
		res := "Reason: " + error.Error(err)
		return res
	}
	r, _ := ioutil.ReadAll(resp.Body)
	res := "Result4: \n---------------------\n" + string(r)
	return res
}
