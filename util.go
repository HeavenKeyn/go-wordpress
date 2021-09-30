package wordpress

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func getResponse(url string, v interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	return unmarshalResponse(resp, v)
}

func unmarshalResponse(resp *http.Response, v interface{}) error {
	defer func() {
		_ = resp.Body.Close()
	}()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}
