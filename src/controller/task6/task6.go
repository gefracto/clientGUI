package task6

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/clientGUI/src/model"
)

func T6(length, maxsquare string) string {
	var data model.T6
	var err error
	var body []byte
	var resp *http.Response
	data.Length, _ = strconv.Atoi(length)
	data.MaxSquare, _ = strconv.Atoi(maxsquare)

	if body, err = json.Marshal(data); err != nil {
		return "Ошибка маршалинга"
	}

	if resp, err = http.Post("http://localhost:1111/task/6", "application/json", bytes.NewBuffer(body)); err != nil {
		res := "Reason: " + error.Error(err)
		return res
	}
	r, _ := ioutil.ReadAll(resp.Body)
	res := "Result6: \n---------------------\n" + string(r)
	return res
}
