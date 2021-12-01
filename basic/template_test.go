package basic

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"testing"
)

func renderScript(script string, data map[string]interface{}) (result []byte, err error) {
	t := template.New("test")

	tpl, err := t.Parse(script)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	resBytes := bytes.NewBufferString("")
	if err = tpl.Execute(resBytes, data); err != nil {
		fmt.Println(err.Error())
		return
	}

	result, _ = ioutil.ReadAll(resBytes)

	return
}

func Test_renderScript(t *testing.T) {
	str := `{{.aa}}
{{.bb}}`

	data := map[string]interface{}{
		"aa": "hello\nworld",
		"bb": []string{"hello"},
	}

	result, err := renderScript(str, data)
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Log(string(result))
}

func Test_renderHyphen(t *testing.T) {
	str := `{{ index .hello "ph-xxx-dddd" }}`

	data := map[string]interface{}{
		"hello": map[string]interface{}{
			"ph-xxx-dddd": "143432"},
	}

	result, err := renderScript(str, data)
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Log(string(result))

}
