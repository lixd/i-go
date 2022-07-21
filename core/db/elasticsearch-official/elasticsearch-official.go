package elasticsearch

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"i-go/utils"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/estransport"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

/*
official sdk
github.com/elastic/go-elasticsearch

感觉不是很好用,日后再说

*/
var ESClientOfficial *elasticsearch.Client

type ESConf struct {
	Address  []string      `json:"address"`
	Username string        `json:"username"`
	Password string        `json:"password"`
	CertFile string        `json:"certFile"` //  "/tmp/test-certs/test-name-1.pem"
	Timeout  time.Duration `json:"timeout"`
}

func Init() {
	defer utils.InitLog("ElasticSearch")()

	c, err := parseConf()
	if err != nil {
		panic(err)
	}
	ESClientOfficial, err = newClient(c)
	if err != nil {
		panic(err)
	}
	res, err := ESClientOfficial.Info()
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	if res.IsError() {
		logrus.Fatalf("Error: %s", res.String())
	}
	// Deserialize the response into a map.
	r := make(map[string]interface{})
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		logrus.Fatalf("Error parsing the response body: %s", err)
	}
	// Print client and server version numbers.
	logrus.Printf("Client: %s", elasticsearch.Version)
	logrus.Printf("Server: %s", r["version"].(map[string]interface{})["number"])
	logrus.Println(strings.Repeat("~", 37))
	fmt.Printf("elasticsearch info %s\n", res)

}

func newClient(c *ESConf) (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: c.Address,
	}
	if c.Username != "" {
		cfg.Username = c.Username
		cfg.Password = c.Password
	}
	// TLS
	if c.CertFile != "" {
		cert, _ := ioutil.ReadFile(c.CertFile)
		cfg.CACert = cert
	}
	if c.Timeout != 0 {
		cfg.Transport = &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS11,
				// ...
			},
			// ...
		}
	}
	cfg.Logger = &estransport.JSONLogger{}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return &elasticsearch.Client{}, err
	}
	return client, nil
}

func parseConf() (*ESConf, error) {
	var c ESConf
	if err := viper.UnmarshalKey("elasticsearch-official", &c); err != nil {
		return &ESConf{}, err
	}
	if len(c.Address) == 0 {
		return &ESConf{}, errors.New("elasticsearch conf nil")
	}
	return &c, nil
}
