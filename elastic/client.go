package elastic

import (
	"bytes"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"strings"
)

const esHosts = "http://10.20.97.39:9200"

var globalEsClient *EsClient

type EsClient struct {
	client *elasticsearch.Client
}

func InitEs() error {
	if globalEsClient != nil {
		return nil
	}

	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: strings.Split(esHosts, ","),
	})
	if err != nil {
		return err
	}

	globalEsClient = &EsClient{
		client: es,
	}
	return nil
}

type Idc struct {
	Id             int    `json:"id" auto:"true"`
	Name           string `json:"name"`
	ItopId         int    `json:"itop_id"`
	Location       string `json:"location"`
	LoginProxy     string `json:"login_proxy"`
	IdcSocketProxy string `json:"idc_socket_proxy"`
	IdcHttpProxy   string `json:"idc_http_proxy"`
	IsCooperation  bool   `json:"is_cooperation"`
}

// index like table
// docs like  rows
func (es *EsClient) WriterData(index string, data []byte) (affected int64, err error) {

	rsp, err := es.client.Create(index, "", bytes.NewReader(data))
	if err != nil {
		return
	}

	if rsp.StatusCode != 200 {
		err = fmt.Errorf(strings.Join(rsp.Warnings(), "\n"))
		return
	}

	return
}
