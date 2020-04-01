package extra

import "testing"

func Test_download(t *testing.T) {
	if err := DownLoad("http://idcos.com/"); err != nil {
		t.Error(err)
	}
}

func TestReadAndWrite(t *testing.T) {
	ReadAndWrite()
}
