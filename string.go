package basic

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/url"
)

func Substr(str string, start, end int) string {
	if len(str) < end {
		end = len(str)
	}
	return string([]rune(str)[start:end])
}

func LnTest(str string) {
	buffer := bytes.NewBufferString("")

	fmt.Fprintf(buffer, "%s\n", "hello")
	//fmt.Fprintf(buffer, "%s", str)
	fmt.Fprintln(buffer, str)
	fmt.Println(buffer.String())
}

func EncodeStr() {
	v := url.Values{}
	v.Add("name", "2018-10-11 10:00:00")
	v.Add("age", "10")

	fmt.Println(v.Encode())
}

func Base64Encode(str string) {
	encode := base64.StdEncoding.EncodeToString([]byte(str))
	fmt.Printf("encode str %s is %s\n", str, encode)
	unicode, _ := base64.StdEncoding.DecodeString(encode)
	fmt.Printf("decode str %s is %s\n ", encode, string(unicode))
}

