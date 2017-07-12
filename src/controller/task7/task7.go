package task7

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/clientGUI/src/model"
)

func T7(FILE string) string {
	var data model.T7
	var err error
	var body []byte
	var resp *http.Response
	data.File = FILE
	if body, err = json.Marshal(data); err != nil {
		return "Ошибка маршалинга"
	}

	if resp, err = http.Post("http://localhost:1111/task/7", "application/json", bytes.NewBuffer(body)); err != nil {
		res := "Reason: " + error.Error(err)
		return res
	}
	r, _ := ioutil.ReadAll(resp.Body)
	res := "Result7: \n---------------------\n" + string(r)
	return res
}
