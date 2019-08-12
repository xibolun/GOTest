package basic

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/go-resty/resty"
)

func Test_http(t *testing.T) {

	//url

	client := http.DefaultClient

	//resp, err := client.Get("http://10.0.8.6:8086/")
	resp, err := client.Post("http://10.0.8.6:8086/", "application/json", nil)

	if err != nil {
		t.Error(err.Error())
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("read resp body fail, %s", err.Error())
		return
	}

	t.Logf("resp body is %s", string(body))
}

func Test_Resty(t *testing.T) {
	startTime := time.Now()
	defaultCli := resty.New().SetHostURL("http://10.0.8.6").SetTimeout(10 * time.Second).SetRetryMaxWaitTime(10 * time.Second).SetRetryCount(3)
	resp, err := defaultCli.R().Post("http://10.0.8.6:8086")
	fmt.Printf("%s", time.Since(startTime).String())
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Logf("%v", resp)
}
