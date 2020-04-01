package basic

import (
	"encoding/json"
	"testing"
	"fmt"
)

func TestJsonArray2Struct(t *testing.T) {

	str := `
     {
  "cc": [
    "string"
  ],
  "content": "string",
  "html": true,
  "receivers": [
    "string"
  ],
  "subject": "string"
}`

	JsonObj2Struct(str, "SendEmailForm")

}

func TestStrConvert(t *testing.T) {
	fmt.Println(FirstUpper("hell"))
}

// RespBody 响应结构体
type RespBody struct {
	Status  string                 `json:"status"`
	Message string                 `json:"message"`
	Content map[string]interface{} `json:"content,omitempty"`
}

func TestUnmashal(t *testing.T) {
	str := `  {"status":"success","message":"操作成功","content":{"content":[{"sn":"fajdlfjalf","app_user_name":"","app_user":"","app_user_uuid":"","company":"","model_name":"","buyer_name":"","buy_time":"","expired_time":"","power_supply_num":0,"assign_status":"","oob_username":"","oob_password":"","level":0,"usage":"","environment":"","spec":"","asset_number":"","location":"","location_id":0,"biz_ip":"","oob_ip":"","os":"","status":"","opt_username":"","opt_user":"","opt_user_uuid":"","user_id":"","remark":"","origin_node":"","content":""}],"import_result":{"limit":8000,"repeat":0,"import_num":1,"total_now":515},"import_status":true,"message":"操作成功","record_count":1,"status":"success"}}`

	var out RespBody

	if err := json.Unmarshal([]byte(str), &out); err != nil {
		t.Error(err)
		return
	}

	fmt.Println("done")

}

func TestJson(t *testing.T) {
	mapper := make(map[string]interface{})
	mapper["age"] = 10

	fmt.Println(ToJsonString(mapper))

	var array []map[string]interface{}

	fmt.Println(ToJsonString(array))

}

// progressLog 进度日志
type progressLog struct {
	IsOk     bool    `json:"is_ok"`
	Progress float64 `json:"progress"`
	Log      string  `json:"log"`
	Title    string  `json:"title"`
}

func TestProgressLog(t *testing.T) {
	progressConf := `{"mount":[{"is_ok":false,"progress":-1,"log":"mount iso fail","title":"挂载镜像失败"},{"is_ok":true,"progress":0.2,"log":"mount iso success","title":"挂载镜像成功(20%)"}],"setBoot":[{"is_ok":false,"progress":-1,"log":"set boot from iso fail","title":"镜像重启失败"},{"is_ok":true,"progress":0.3,"log":"set boot from iso success","title":"镜像重启成功(30%)"}],"powerOn":[{"is_ok":false,"progress":0.9,"log":"oob os installation success, power status is off","title":"装机完成(90%)"}],"restart":[{"is_ok":false,"progress":-1,"log":"reboot device fail","title":"设备重启失败"},{"is_ok":true,"progress":0.99,"log":"reboot device success","title":"设备重启完成(99%)"}],"unmount":[{"is_ok":true,"progress":1,"log":"unmount iso  success","title":"卸载镜像成功(100%)"}]}`
	var conf map[string][]*progressLog
	_ = json.Unmarshal([]byte(progressConf), &conf)
	fmt.Println(conf)
}

func TestProgressLogString(t *testing.T) {
	var logs []*progressLog
	process1 := &progressLog{
		IsOk:     true,
		Progress: -1,
		Log:      "hello",
		Title:    "title",
	}
	process2 := &progressLog{
		IsOk:     true,
		Progress: -1,
		Log:      "hello",
		Title:    "title",
	}

	logs = append(logs, process1)
	logs = append(logs, process2)

	mapper := make(map[string][]*progressLog)
	mapper["mount"] = logs

	fmt.Println(ToJsonString(mapper))

}

func TestUnmashal2(t *testing.T) {
	// null 和 0都會轉成0
	str := `{"age": 0}`
	type Age struct {
		Age int `json:"age,omitempty"`
	}
	var age Age

	json.Unmarshal([]byte(str), &age)

	fmt.Println(age.Age)
}

func TestUnmarshalOption(t *testing.T) {
	options := "{\"username\":\"salt-api\", \"password\":\"yunjikeji\"}"
	//ChannelOption salt channel option
	type ChannelOption struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var option ChannelOption
	err := json.Unmarshal([]byte(options), &option)
	if err != nil {
		t.Error(err)
	}

}
