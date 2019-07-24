package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/http/httputil"
)

func main() {
	request, err := http.NewRequest(http.MethodGet, "https://www.lixueduan.com", nil)
	if err != nil {
		logrus.Errorf("http.NewRequest err=%v", err)
	}
	request.Header.Add("user-agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36"+
		" (KHTML, like Gecko) Chrome/72.0.3626.121 Mobile Safari/537.36")
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println(req)
			return nil
		},
	}
	response, err := client.Do(request)
	if err != nil {
		logrus.Errorf("http.DefaultClient.Do(request) err=%v", err)
	}

	defer response.Body.Close()

	bytes, err := httputil.DumpResponse(response, true)
	if err != nil {
		logrus.Errorf("httputil.DumpResponse err=%v", err)
	}

	fmt.Printf("resp =%s \n", bytes)
}
