package smartctl

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"os/exec"
	"runtime"
	"strings"

	"github.com/qbox/jarvis/common/collection/exector"
	"github.com/qbox/jarvis/common/utils"

	"github.com/tidwall/gjson"
	"github.com/zeromicro/go-zero/core/logx"
)

const (
	smartctlCmdScan = "-c smartctl --scan-open -j"
	smartctlCmdInfo = "-c smartctl -s on  -i -j %s"
	smartctlCmdAH   = "-c smartctl -A -H -d %s -j %s" // smartctl -A -H -d sat /dev/sda -j
)

var (
	SatMetricRules = []*MetricRule{
		{Key: "Reallocated_Sector_Ct", Type: RuleTypeLt, Threshold: 50},        // 重分配扇区计数
		{Key: "Reallocation_Event_Ct", Type: RuleTypeLt, Threshold: 50},        // 重分配扇区计数
		{Key: "Temperature_Case", Type: RuleTypeGt, Threshold: 60},             // 磁盘温度
		{Key: "Spin_Retry_Count", Type: RuleTypeGt, Threshold: 10},             // 电机起转重试
		{Key: "Read_Error_Rate", Type: RuleTypeGt, Threshold: 10},              // 底层数据读取错误率
		{Key: "Soft_Read_Error_Rate", Type: RuleTypeGt, Threshold: 10},         // 逻辑读取错误率
		{Key: "Reallocation_Event_Ct", Type: RuleTypeGt, Threshold: 10},        // 终端校验出错（HP专有）
		{Key: "Current_Pending_Sector_Count", Type: RuleTypeGt, Threshold: 10}, // 等候重定的扇区计数
		{Key: "Uncorrectable_Sector_Count", Type: RuleTypeGt, Threshold: 10},   // 无法校正的扇区计数
		{Key: "Temperature_Case", Type: RuleTypeGt, Threshold: 80},             // 磁盘表面温度
		{Key: "Temperature_Internal_raw", Type: RuleTypeGt, Threshold: 80},     // 磁盘表面温度
	}

	ScsiMetricRules = []*MetricRule{
		{Key: "temperature", Type: RuleTypeGt, Threshold: 80}, // 磁盘温度
	}

	NvmeMetricRules = []*MetricRule{
		{Key: "critical_warning", Type: RuleTypeGt, Threshold: 1}, // 错误警告
		{Key: "temperature", Type: RuleTypeGt, Threshold: 80},     // 磁盘温度
		{Key: "percentage_used", Type: RuleTypeGt, Threshold: 90}, // 使用率
		{Key: "media_errors", Type: RuleTypeGt, Threshold: 1},     // 错误信息
	}
)

type DiskDevice struct {
	NodeId          string                 `json:"nodeId"`
	Name            string                 `json:"name"`
	InfoName        string                 `json:"info_name"`
	Type            string                 `json:"type"`
	Protocol        string                 `json:"protocol"`
	ModelFamily     string                 `json:"model_family"`
	ModelName       string                 `json:"model_name"`
	SerialNumber    string                 `json:"serial_number"`
	FirmwareVersion string                 `json:"firmware_version"`
	Attributes      map[string]interface{} `json:"attributes"`
}

type DiskChecker struct {
	log      logx.Logger
	executor exector.Executor
}

func NewDiskChecker() *DiskChecker {
	return &DiskChecker{
		log:      logx.WithContext(context.Background()),
		executor: exector.DefaultExecutor,
	}

}

func (c *DiskChecker) preCheck() bool {
	if runtime.GOOS == "darwin" {
		return true
	}
	path, err := exec.LookPath("smartctl")
	return err != nil && path != ""
}

func (c *DiskChecker) checkResultStatus(output []byte) bool {
	r := gjson.Get(string(output), "smartctl.exit_status")
	zeroExit := r.Type == gjson.Number && r.Int() == 0 && r.Exists()

	ispass := gjson.Get(string(output), "smart_status.passed")
	if !ispass.Exists() {
		return zeroExit
	}

	return zeroExit && ispass.IsBool() && ispass.Bool()
}

func (c *DiskChecker) scan() (devices []*DiskDevice, err error) {
	output, err := c.executor.Exec(nil, "/bin/bash", strings.Split(smartctlCmdScan, " ")...)
	if err != nil {
		c.log.Errorf("fail exec cmd %s, %s", smartctlCmdScan, err.Error())
		return
	}

	if !c.checkResultStatus(output) {
		err = fmt.Errorf("cmd exec is not zero %s, %s", smartctlCmdScan, string(output))
		c.log.Error(err)
		return
	}

	r := gjson.Get(string(output), "devices")
	devices = make([]*DiskDevice, 0)
	_ = json.Unmarshal([]byte(r.String()), &devices)
	return
}

func (c *DiskChecker) FillDeviceInfo(device *DiskDevice) (err error) {
	cmd := fmt.Sprintf(smartctlCmdInfo, device.Name)
	output, err := c.executor.Exec(nil, "/bin/bash", strings.Split(cmd, " ")...)
	if err != nil {
		c.log.Errorf("fail exec cmd %s, %s", cmd, err.Error())
		return
	}

	if !c.checkResultStatus(output) {
		err = fmt.Errorf("cmd exec is not zero %s, %s", cmd, string(output))
		c.log.Error(err)
		return
	}

	device.ModelName = gjson.Get(string(output), "model_name").String()
	device.SerialNumber = gjson.Get(string(output), "serial_number").String()
	device.ModelFamily = gjson.Get(string(output), "model_family").String()
	device.FirmwareVersion = gjson.Get(string(output), "firmware_version").String()

	if runtime.GOOS == "darwin" {
		device.NodeId = uuid.NewString()
	} else {
		device.NodeId = string(utils.ReadFileIgnoreErr("/etc/machine-id"))
	}

	return
}

func (c *DiskChecker) FillSatMetrics(device *DiskDevice) (err error) {
	cmd := fmt.Sprintf(smartctlCmdAH, device.Type, device.Name)
	output, err := c.executor.Exec(nil, "/bin/bash", strings.Split(cmd, " ")...)
	if err != nil {
		c.log.Errorf("fail exec cmd %s, %s", cmd, err.Error())
		return
	}

	if !c.checkResultStatus(output) {
		err = fmt.Errorf("cmd exec is not zero %s, %s", cmd, string(output))
		c.log.Error(err)
		return
	}

	device.Attributes = make(map[string]interface{})
	tables := gjson.Get(string(output), "ata_smart_attributes.table")
	for _, item := range tables.Array() {
		device.Attributes[item.Get("name").String()] = item.Get("value").String()
	}
	return
}

func (c *DiskChecker) FillNVMEMetrics(device *DiskDevice) (err error) {
	cmd := fmt.Sprintf(smartctlCmdAH, device.Type, device.Name)
	output, err := c.executor.Exec(nil, "/bin/bash", strings.Split(cmd, " ")...)
	if err != nil {
		c.log.Errorf("fail exec cmd %s, %s", cmd, err.Error())
		return
	}

	if !c.checkResultStatus(output) {
		err = fmt.Errorf("cmd exec is not zero %s, %s", cmd, string(output))
		c.log.Error(err)
		return
	}

	device.Attributes = make(map[string]interface{})
	err = json.Unmarshal([]byte(gjson.Get(string(output), "nvme_smart_health_information_log").String()), &device.Attributes)
	if err != nil {
		c.log.Errorf("fail unmarshal output, %s", string(output))
		return
	}

	return
}
func (c *DiskChecker) FillSCSIMetrics(device *DiskDevice) (err error) {
	//https://linkcloud-admin.qiniu.io/fullpage/ssh/14ebb440448927a9b6c39b1bf8e1eb29 此节点有scsi盘
	cmd := fmt.Sprintf(smartctlCmdAH, device.Type, device.Name)
	output, err := c.executor.Exec(nil, "/bin/bash", strings.Split(cmd, " ")...)
	if err != nil {
		c.log.Errorf("fail exec cmd %s, %s", cmd, err.Error())
		return
	}

	if !c.checkResultStatus(output) {
		err = fmt.Errorf("cmd exec is not zero %s, %s", cmd, string(output))
		c.log.Error(err)
		return
	}

	device.Attributes = make(map[string]interface{})
	device.Attributes["temperature"] = gjson.Get(string(output), "temperature.current").String()

	return
}

func (c *DiskChecker) collect() (err error) {
	if !c.preCheck() {
		return fmt.Errorf("disk pre check fail")
	}

	items, err := c.scan()
	if err != nil {
		return
	}

	for _, item := range items {
		switch strings.ToLower(item.Type) {
		case "sat", "usbjmicron", "usbprolific", "usbsunplus":
			_ = c.FillSatMetrics(item)
			c.loggingMetric(item, SatMetricRules)
		case "nvme", "sntasmedia", "sntjmicron", "sntrealtek":
			_ = c.FillNVMEMetrics(item)
			c.loggingMetric(item, NvmeMetricRules)
		case "scsi":
			_ = c.FillSCSIMetrics(item)
			c.loggingMetric(item, ScsiMetricRules)
		default:
			c.log.Errorf("unknown disk type, %s", item.Type)
		}
	}
	return
}

func (c *DiskChecker) loggingMetric(disk *DiskDevice, rules []*MetricRule) {
	for _, item := range rules {
		targetV, ok := disk.Attributes[item.Key]
		if !ok {
			continue
		}
		if item.IsPass(targetV) {
			continue
		}
		fmt.Println(fmt.Sprintf(`jarvis_disk_%s_error{nodeId="%s",device="%s"} %v`, strings.ToLower(item.Key),
			disk.NodeId, disk.Name, targetV))
	}

}

func main() {
	checker := NewDiskChecker()
	_ = checker.collect()
}
