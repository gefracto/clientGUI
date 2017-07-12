package task3

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/clientGUI/src/model"
)

func T3(names, a, b, c []string) string {
	var data []model.Triangle
	var err error
	var body []byte
	var resp *http.Response
	for i, v := range names {
		a_, _ := strconv.ParseFloat(a[i], 64)
		b_, _ := strconv.ParseFloat(b[i], 64)
		c_, _ := strconv.ParseFloat(c[i], 64)
		data = append(data, model.Triangle{v, a_, b_, c_})
	}
	if body, err = json.Marshal(data); err != nil {
		return "Ошибка маршалинга"
	}
	if resp, err = http.Post("http://localhost:1111/task/3", "application/json", bytes.NewBuffer(body)); err != nil {
		res := "Reason: " + error.Error(err)
		return res
	}
	r, _ := ioutil.ReadAll(resp.Body)
	res := "Result3: \n---------------------\n" + string(r)
	return res
}
