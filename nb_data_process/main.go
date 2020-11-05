package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// 用于处理宁波银行因服务重启，导致act2_host表的数据丢失的问题
// 连接数据库配置
// export DB_HOST=xxxx
// export DB_USER=root
// export DB_PASSWORD=xxx
// export DB_NAME=cloud-act2
// export REGISTER_PARAM_PATH=/tmp/regParam.json
func main() {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	regStrPath := os.Getenv("REGISTER_PARAM_PATH")
	bytes, err := ioutil.ReadFile(regStrPath)
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local&timeout=5s&readTimeout=5m&writeTimeout=5m", dbUser, dbPassword, dbHost, dbName))

	if err != nil {
		fmt.Println(err)
	}
	//关闭数据库
	defer db.Close()

	// find all lost hostID ip record
	hostIps := make([]Act2HostIP, 0)
	if err := db.Model(Act2HostIP{}.TableName()).Raw(`SELECT * FROM act2_host_ip WHERE host_id not in (select id from act2_host)`).Scan(&hostIps).Error; err != nil {
		panic(err)
	}

	// lost hostID ip record
	dbIPMap := recordToIPMap(hostIps)

	for _, regStr := range strings.Split(string(bytes), "\n") {
		// register data
		regParam := parseRegParam(regStr)
		regIPMap := regParamToIPMap(regParam)

		// get proxy from register data proxy_server
		proxy := getProxy(regParam.Master.Server, db)

		// 根据IP去register data里面找到对应的host主机信息；
		// idc_id, entity_id,add_time,status,os_type,minion_version,proxy_id
		hosts := make([]*Act2Host, 0)
		for ip, v := range dbIPMap {
			if minion, ok := regIPMap[ip]; ok {
				hosts = append(hosts, &Act2Host{
					ID:            v.HostID,
					IdcID:         proxy.IdcID,
					EntityID:      minion.Sn,
					ProxyID:       proxy.ID,
					AddTime:       v.AddTime,
					OsType:        minion.OsType,
					Status:        minion.Status,
					MinionVersion: minion.MinionVersion,
				})
			}
		}

		for i := range hosts {
			fmt.Printf("INSERT INTO act2_host (id, idc_id, entity_id, add_time, status, os_type, minion_version, proxy_id) VALUES ('%s','%s','%s','%s','%s','%s','%s','%s');\n", hosts[i].ID, hosts[i].IdcID, hosts[i].EntityID, hosts[i].AddTime.Format("2006-01-02 15:04:03"), hosts[i].Status, hosts[i].OsType, hosts[i].MinionVersion, hosts[i].ProxyID)
		}
	}
}

// recordToIPMap
// 已经确保IP不会重复，所以可以拿IP来做为K
// select * from (SELECT ip,count(1) as count FROM act2_host_ip WHERE host_id not in (select id from act2_host) GROUP BY ip) as tmp WHERE tmp.count >1;
func recordToIPMap(hostIps []Act2HostIP) map[string]*Act2HostIP {
	ipMap := make(map[string]*Act2HostIP)
	for i := range hostIps {
		ipMap[hostIps[i].IP] = &hostIps[i]
	}
	return ipMap
}
func regParamToIPMap(param *RegParam) map[string]*Minion {
	minions := param.Minion

	ipMap := make(map[string]*Minion)
	for i := range minions {
		for _, ip := range minions[i].IPs {
			if strings.Contains(ip, "172.") {
				continue
			}

			ipMap[ip] = &minions[i]
		}
	}
	return ipMap
}

func getProxy(proxysrv string, db *gorm.DB) *Act2Proxy {
	proxy := Act2Proxy{}
	db.Model(Act2Proxy{}.TableName()).Where("server = ?", proxysrv).Find(&proxy)
	return &proxy
}

// Register 上报数据传输对象
type RegParam struct {
	Master Master   `json:"master"`
	Minion []Minion `json:"minions"`
}
type Master struct {
	Sn      string      `json:"sn"`
	Server  string      `json:"server"`
	Status  string      `json:"status"`
	Idc     string      `json:"idc"`
	Type    string      `json:"type"`
	Options Options     `json:"options"`
	HostMax HostMaxInfo `json:"hostMax"`
}

//HostMaxInfo 主机数量限制
type HostMaxInfo struct {
	Salt    int `json:"salt"`
	SSH     int `json:"ssh"`
	Puppet  int `json:"puppet"`
	Ansible int `json:"ansible"`
}

type Options struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Minion struct {
	Sn            string   `json:"sn"`
	IPs           []string `json:"ips"`
	Status        string   `json:"status"`
	OsType        string   `json:"os_type"`
	MinionVersion string   `json:"minionVersion"`
}

type Act2HostIP struct {
	ID      string    `gorm:"column:id;primary_key" json:"id"`
	HostID  string    `gorm:"column:host_id" json:"hostId"`
	IP      string    `gorm:"column:ip" json:"ip"`
	AddTime time.Time `gorm:"column:add_time" json:"addTime"`
}

func (Act2HostIP) TableName() string {
	return "act2_host_ip"
}

type Act2Host struct {
	ID            string    `gorm:"column:id;primary_key"`
	IdcID         string    `gorm:"column:idc_id"`
	EntityID      string    `gorm:"column:entity_id"`
	ProxyID       string    `gorm:"column:proxy_id"`
	AddTime       time.Time `gorm:"column:add_time"`
	OsType        string    `gorm:"column:os_type"`
	Status        string    `gorm:"column:status"`
	MinionVersion string    `gorm:"column:minion_version"`
}

func (Act2Host) TableName() string {
	return "act2_host"
}

// Act2Proxy table struct
type Act2Proxy struct {
	ID        string    `gorm:"column:id;primary_key"`
	TwiceTime time.Time `gorm:"column:twice_time"`
	LastTime  time.Time `gorm:"column:last_time"`
	IdcID     string    `gorm:"column:idc_id"`
	Server    string    `gorm:"column:server"`
	Type      string    `gorm:"column:type"`
	Status    string    `gorm:"column:status"`
	Options   string    `gorm:"column:options"`
	HostMax   string    `gorm:"column:host_max"`
}

// TableName convert
func (Act2Proxy) TableName() string {
	return "act2_proxy"
}

func parseRegParam(regStr string) *RegParam {

	regParam := RegParam{}
	_ = json.Unmarshal([]byte(regStr), &regParam)

	return &regParam
}
