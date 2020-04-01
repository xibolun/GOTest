package extra

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

/**
用golang实现一个下面的命令
	curl http://192.168.1.27/tmp.tgz  | tar -zxpf  - -C /tmp/test
*/
var cache = make([]byte, 102400)

func DownLoad(url string) error {
	//r, w := io.Pipe()
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	_, _ = resp.Body.Read(cache)

	fmt.Println(string(cache))
	return nil
}

func ReadAndWrite() {
	r, w := io.Pipe()

	//load a file
	cache := make([]byte, 4*1024)
	file, _ := os.Open("/tmp/adobegc.log")

	count := 0
	for {
		count++
		_, err := file.Read(cache)
		if err == io.EOF {
			break
		}

		time.Sleep(1 * time.Second)
		//fmt.Println(string(cache))

		if _, err = bytes.NewReader(cache).WriteTo(w); err != nil {
			fmt.Println(err.Error())
		}

	}

	bytes := make([]byte, 4*1024)
	_, _ = r.Read(bytes)
	fmt.Println(string(bytes))
}
