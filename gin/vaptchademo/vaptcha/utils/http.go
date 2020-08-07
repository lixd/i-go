package utils

import (
	"io/ioutil"
	"net/http"
)

func HttpGet(url string) ([]byte, int) {
	var (
		client  = &http.Client{}
		request *http.Request
		resp    *http.Response
		err     error
	)
	request, _ = http.NewRequest(http.MethodGet, url, nil)
	resp, err = client.Do(request)
	if err != nil {
		return nil, http.StatusNotFound
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body, resp.StatusCode
}
