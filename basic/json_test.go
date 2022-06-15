package basic

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
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

func TestUnmarshal(t *testing.T) {
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

func TestUnmarshal2(t *testing.T) {
	// null 和 0都會轉成0
	str := `{"age": 0,"address":[]}`
	type Age struct {
		Age     int      `json:"age,omitempty"`
		Address []string `json:"address,omitempty"`
	}
	var age Age

	json.Unmarshal([]byte(str), &age)

	fmt.Println(age.Age)
}

func TestUnmarshal3(t *testing.T) {
	b := []byte{123, 125}
	var p *progressLog
	if err := json.Unmarshal(b, p); err != nil {
		t.Error(err)
	}
	fmt.Printf("%v", p)
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

func TestUnmarshalDiskDriver(t *testing.T) {

	// PhysicalDrive 物理驱动器条目
	type PhysicalDrive struct {
		Name            string `json:"name" comment:"名称"`    // 硬件位置/名称
		Manufacturer    string `json:"vendor" comment:"厂商"`  // 厂商
		Model           string `json:"model" comment:"型号"`   // 型号
		WWN             string `json:"wwn" comment:"WWN"`    // wwn
		SerialNumber    string `json:"serial" comment:"序列号"` // 序列号
		MediaType       string `json:"rota" comment:"媒体类型"`  // 媒体类型
		Size            string `json:"size" comment:"容量"`    // 容量，单位byte。
		FirmwareVersion string `json:"rev" comment:"固件版本"`   // 固件版本
		FirmwareState   string `json:"state" comment:"固件状态"` // 固件状态
	}
	//jsonStr := `{"name": "sda", "rota": "1", "type": "disk", "size": "279.4G", "state": "running", "rev": "1004", "vendor": "TOSHIBA ", "model": "AL13SEB300      ", "serial": "500003971801b808", "wwn": "0x500003971801b808"},`
	jsonStr := ``

	var driver *PhysicalDrive
	if err := json.Unmarshal([]byte(strings.TrimRight(jsonStr, ",")), &driver); err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(driver)
}

func TestDiffJsonArray(t *testing.T) {
	j1 := `[{"ip":"10.10.37.10","interface":"ppp14","in":1559,"out":1643,"uid":0},{"ip":"10.10.16.80","interface":"ppp10","in":840,"out":840,"uid":0},{"ip":"10.10.121.12","interface":"ppp4","in":7028831,"out":115882,"uid":0},{"ip":"10.10.14.254","interface":"ppp0","in":16661381,"out":819008,"uid":0},{"ip":"10.10.31.250","interface":"ppp3","in":2578241,"out":24095,"uid":0},{"ip":"10.10.107.49","interface":"ppp6","in":10498,"out":2297,"uid":0},{"ip":"10.10.36.235","interface":"ppp15","in":1086,"out":1225,"uid":0},{"ip":"10.10.72.14","interface":"ppp2","in":4403,"out":147492,"uid":0},{"ip":"10.10.103.193","interface":"ppp5","in":21790,"out":2575,"uid":0},{"ip":"10.10.60.171","interface":"ppp13","in":1086,"out":1225,"uid":0},{"ip":"10.10.0.103","interface":"ppp9","in":1087,"out":1226,"uid":0},{"ip":"10.10.22.204","interface":"ppp8","in":7903754,"out":55383,"uid":0},{"ip":"10.10.37.4","interface":"ppp1","in":1681,"out":4604,"uid":0},{"ip":"10.10.106.60","interface":"ppp11","in":1484,"out":23727,"uid":0},{"ip":"10.10.112.247","interface":"ppp12","in":2020,"out":2059,"uid":0},{"ip":"10.10.8.100","interface":"ppp7","in":1552,"out":1641,"uid":0}]`
	j2 := `[{"ip":"100.86.15.124","interface":"ppp2","natType":"Full Cone","externalIP":"111.43.27.181","externalPort":"12866"},{"ip":"100.86.15.193","interface":"ppp20","natType":"Blocked","externalIP":"None","externalPort":"None"},{"ip":"fe80::f016:bcff:fe2a:f798","interface":"vth27","natType":"Blocked","externalIP":"None","externalPort":"None"},{"ip":"100.86.15.158","interface":"ppp6","natType":"Blocked","externalIP":"None","externalPort":"None"},{"ip":"fe80::6436:abff:fe4d:df","interface":"vth12","natType":"Blocked","externalIP":"None","externalPort":"None"},{"ip":"100.86.14.251","interface":"ppp17","natType":"Full Cone","externalIP":"111.43.27.128","externalPort":"12859"},{"ip":"100.86.13.135","interface":"ppp29","natType":"","externalIP":"","externalPort":""},{"ip":"100.86.15.65","interface":"ppp23","natType":"Blocked","externalIP":"None","externalPort":"None"},{"ip":"fe80::6e92:bfff:fe4d:301","interface":"eth3","natType":"Blocked","externalIP":"None","externalPort":"None"},{"ip":"fe80::184e:2eff:fedb:a253","interface":"vth25","natType":"Blocked","externalIP":"None","externalPort":"None"},{"ip":"fe80::415:bcff:fed3:5e8b","interface":"vth20","natType":"Blocked","externalIP":"None","externalPort":"None"},{"ip":"fe80::8c03:73ff:feee:5e81","interface":"vth23","natType":"Blocked","externalIP":"None","externalPort":"None"},{"ip":"fe80::8a6:2cff:fe54:bf9d","interface":"vth24","natType":"Blocked","externalIP":"None","externalPort":"None"},{"ip":"100.86.13.105","interface":"ppp13","natType":"","externalIP":"","externalPort":""},{"ip":"100.86.13.85","interface":"ppp18","natType":"Blocked","externalIP":"None","externalPort":"None"},{"ip":"100.86.14.227","interface":"ppp21","natType":"Blocked","externalIP":"None","externalPort":"None"},{"ip":"100.86.14.51","interface":"ppp24","natType":"","externalIP":"","externalPort":""},{"ip":"100.86.15.126","interface":"ppp14","natType":"Blocked","externalIP":"None","externalPort":"None"},{"ip":"100.86.15.97","interface":"ppp12","natType":"Blocked","externalIP":"None","externalPort":"None"},{"ip":"fe80::c03d:c7ff:fe1a:1c8c","interface":"vth7","natType":"Blocked","externalIP":"None","externalPort":"None"},{"ip":"fe80::483c:cfff:feeb:4947","interface":"vth17","natType":"Blocked","externalIP":"None","externalPort":"None"},{"ip":"100.86.12.134","interface":"ppp26","natType":"Full Cone","externalIP":"111.43.27.111","externalPort":"10358"},{"ip":"100.86.12.58","interface":"ppp27","natType":"Blocked","externalIP":"None","externalPort":"None"},{"ip":"fe80::9c55:8ff:feea:409b","interface":"vth14","natType":"Blocked","externalIP":"None","externalPort":"None"},{"ip":"fe80::bc4d:b6ff:feca:d996","interface":"vth16","natType":"Blocked","externalIP":"None","externalPort":"None"},{"ip":"fe80::74b2:6cff:fe37:4a91","interface":"vth5","natType":"Blocked","externalIP":"None","externalPort":"None"},{"ip":"fe80::d402:3dff:fe96:1de8","interface":"vth10","natType":"Blocked","externalIP":"None","externalPort":"None"},{"ip":"100.86.1.18","interface":"ppp25","natType":"Full Cone","externalIP":"111.43.27.143","externalPort":"3686"},{"ip":"100.86.14.190","interface":"ppp10","natType":"Full Cone","externalIP":"111.43.27.143","externalPort":"15705"},{"ip":"fe80::ac89:6fff:fe51:7b02","interface":"vth1","natType":"Blocked","externalIP":"None","externalPort":"None"},{"ip":"fe80::b434:45ff:fec0:399e","interface":"vth18","natType":"Blocked","externalIP":"None","externalPort":"None"},{"ip":"fe80::44ea:f3ff:fe6f:663e","interface":"vth8","natType":"Blocked","externalIP":"None","externalPort":"None"},{"ip":"fe80::a40a:c9ff:fe04:ceaf","interface":"vth0","natType":"Blocked","externalIP":"None","externalPort":"None"},{"ip":"100.86.14.170","interface":"ppp4","natType":"Full Cone","externalIP":"111.43.27.135","externalPort":"6453"},{"ip":"fe80::10ae:1aff:fe82:2f1","interface":"vth15","natType":"Blocked","externalIP":"None","externalPort":"None"},{"ip":"100.86.12.212","interface":"ppp7","natType":"Full Cone","externalIP":"111.43.27.70","externalPort":"14273"},{"ip":"100.86.15.20","interface":"ppp3","natType":"Full Cone","externalIP":"111.43.27.153","externalPort":"8969"},{"ip":"100.86.15.203","interface":"ppp16","natType":"Full Cone","externalIP":"111.43.27.150","externalPort":"18598"},{"ip":"fe80::8013:4bff:fe7d:2ba6","interface":"vth26","natType":"Blocked","externalIP":"None","externalPort":"None"},{"ip":"fe80::3042:87ff:fe97:aec2","interface":"vth22","natType":"Blocked","externalIP":"None","externalPort":"None"},{"ip":"fe80::31:c3ff:fe0a:27ae","interface":"vth13","natType":"Blocked","externalIP":"None","externalPort":"None"},{"ip":"fe80::34e8:e0ff:fe42:30cf","interface":"vth3","natType":"Blocked","externalIP":"None","externalPort":"None"},{"ip":"100.86.12.79","interface":"ppp28","natType":"Blocked","externalIP":"None","externalPort":"None"},{"ip":"100.86.12.98","interface":"ppp22","natType":"Blocked","externalIP":"None","externalPort":"None"},{"ip":"100.86.14.27","interface":"ppp15","natType":"Full Cone","externalIP":"111.43.27.184","externalPort":"7149"},{"ip":"100.86.14.100","interface":"ppp11","natType":"Full Cone","externalIP":"111.43.27.178","externalPort":"3613"},{"ip":"100.86.14.43","interface":"ppp1","natType":"Blocked","externalIP":"None","externalPort":"None"},{"ip":"fe80::209d:13ff:feb3:ea11","interface":"vth11","natType":"Blocked","externalIP":"None","externalPort":"None"},{"ip":"100.86.14.70","interface":"ppp19","natType":"Full Cone","externalIP":"111.43.27.135","externalPort":"10707"},{"ip":"100.86.14.42","interface":"ppp9","natType":"Full Cone","externalIP":"111.43.27.188","externalPort":"10265"},{"ip":"fe80::c05c:26ff:fea9:6d9e","interface":"vth6","natType":"Blocked","externalIP":"None","externalPort":"None"},{"ip":"fe80::9c1f:65ff:fe94:954d","interface":"vth29","natType":"Blocked","externalIP":"None","externalPort":"None"},{"ip":"100.86.14.47","interface":"ppp8","natType":"","externalIP":"","externalPort":""},{"ip":"192.168.1.14","interface":"eth3","natType":"Blocked","externalIP":"None","externalPort":"None"},{"ip":"100.86.14.95","interface":"ppp0","natType":"Full Cone","externalIP":"111.43.27.146","externalPort":"4084"},{"ip":"fe80::505c:83ff:fe65:5ecf","interface":"vth2","natType":"Blocked","externalIP":"None","externalPort":"None"},{"ip":"fe80::10fd:9aff:fe12:8458","interface":"vth28","natType":"Blocked","externalIP":"None","externalPort":"None"},{"ip":"100.86.14.88","interface":"ppp5","natType":"Blocked","externalIP":"None","externalPort":"None"},{"ip":"fe80::d4c6:ccff:fee0:7be5","interface":"vth21","natType":"Blocked","externalIP":"None","externalPort":"None"},{"ip":"fe80::ac76:eaff:fef4:775d","interface":"vth19","natType":"Blocked","externalIP":"None","externalPort":"None"},{"ip":"fe80::e07b:18ff:fe80:21a6","interface":"vth4","natType":"Blocked","externalIP":"None","externalPort":"None"},{"ip":"fe80::d4b8:edff:fec8:1a7d","interface":"vth9","natType":"Blocked","externalIP":"None","externalPort":"None"}]`

	pk := []string{"ip", "interface"}

	m1 := make([]map[string]string, 0)
	m2 := make([]map[string]string, 0)

	_ = json.Unmarshal([]byte(j1), &m1)
	_ = json.Unmarshal([]byte(j2), &m2)

	arrayToMap := func(array []map[string]string) map[string]map[string]string {
		ret := make(map[string]map[string]string)
		for i := range array {
			pkV := ""
			for _, item := range pk {
				pkV += array[i][item]
			}
			ret[pkV] = array[i]
		}
		return ret
	}

	v1 := arrayToMap(m1)
	v2 := arrayToMap(m2)

	for k1 := range v1 {
		if _, ok := v2[k1]; !ok {
			t.Logf("j2 do not have %s", k1)
		}
	}

	for k2 := range v2 {
		if _, ok := v1[k2]; !ok {
			t.Logf("j1 do not have %s", k2)
		}
	}

}

func TestAA(t *testing.T) {
	str := "hello world\n"
	fmt.Printf("%x", md5.Sum([]byte(str)))

}
