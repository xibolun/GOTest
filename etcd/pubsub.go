package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/pkg/transport"
	"time"
)

type EPS struct {
	c *clientv3.Client
}

const (
	Queue = "hammer"
	CAFile   = "/Users/pgy/Qiniu/etc/etcd/ca.pem"
	KeyFile  = "/Users/pgy/Qiniu/etc/etcd/client-key.pem"
	CertFile = "/Users/pgy/Qiniu/etc/etcd/client.pem"
	//CAFile   = ""
	//KeyFile  = ""
	//CertFile = ""
)

var gEPS *EPS

func init() {
	if gEPS != nil {
		return
	}

	conf := clientv3.Config{
		Endpoints:   []string{"10.20.97.32:2379"},
		DialTimeout: 5 * time.Second,
	}

	tlsInfo := transport.TLSInfo{
		CertFile: CertFile,
		KeyFile:  KeyFile,
		CAFile:   CAFile,
		//ClientCertAuth: true,
	}
	tlsConfig, err := tlsInfo.ClientConfig()
	if err != nil {
		panic(err)
	}

	conf.TLS = tlsConfig

	eps, err := clientv3.New(conf)
	if err != nil {
		panic(err)
	}

	gEPS = &EPS{c: eps}
}

func ToJsonString(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func Pub(queue string, v interface{}) (err error) {
	fmt.Printf("start to pub into %s\n", queue)
	timeout, _ := context.WithTimeout(context.Background(), 10*time.Second)
	rsp, err := gEPS.c.Put(timeout, queue, ToJsonString(v))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(rsp.OpResponse())
	return
}

func Sub(queue string) (err error) {
	for {
		v := gEPS.c.Watch(context.Background(), queue)
		for rsp := range v {
			err = rsp.Err()
			if err != nil {
				fmt.Println(err.Error())
			}
			for _, ev := range rsp.Events {
				fmt.Println(ToJsonString(ev))
			}

		}
	}
}
