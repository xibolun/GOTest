package boot_lic

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os/exec"
)

// Collector 设备指纹采集器
type Collector interface {
	// Get 返回SHA256哈希算法计算后的设备指纹
	Get() (fingerprint string)
}

// Get 使用默认采集器采集设备信息指纹
func Get() (fingerprint string) {
	return defaultC.Get()
}

// defaultC 默认采集器
var defaultC = new(LinuxCollector)

// LinuxCollector Linux操作系统下的采集器
type LinuxCollector struct {
}

// Get 返回SHA256哈希算法计算后的设备指纹
func (c *LinuxCollector) Get() (fingerprint string) {
	bs := []byte(fmt.Sprintf("os:%s|kernel:%s|cpu:%s|ip:%s|mac:%s|version:%s|product_name:%s|product_serial:%s|product_uuid:%s|bios_vendor:%s|bios_version:%s|bios_date:%s",
		c.os(),
		c.kernel(),
		c.cpu(),
		c.ip(),
		c.mac(),
		c.vendor(),
		c.productName(),
		c.productSerial(),
		c.productUUID(),
		c.biosVendor(),
		c.biosVersion(),
		c.biosDate(),
	))
	return fmt.Sprintf("%x", sha256.Sum256(bs))
}

func (c *LinuxCollector) os() string {
	bs, _ := ioutil.ReadFile("/etc/redhat-release")
	return base64.StdEncoding.EncodeToString(bs)
}

func (c *LinuxCollector) kernel() string {
	bs, _ := ioutil.ReadFile("/proc/version")
	return base64.StdEncoding.EncodeToString(bs)
}

func (c *LinuxCollector) cpu() string {
	bs, _ := exec.Command("bash", "-c", "awk -F':' '/^model name/ { print $NF; exit }' /proc/cpuinfo").Output()
	return base64.StdEncoding.EncodeToString(bs)
}

func (c *LinuxCollector) ip() string {
	bs, _ := exec.Command("bash", "-c", "ip -o -4 route get 8.8.8.8 | awk '{ print $7 }'").Output()
	return base64.StdEncoding.EncodeToString(bs)
}

func (c *LinuxCollector) mac() string {
	bs, _ := exec.Command("bash", "-c", "cat /sys/class/net/$(ip -o -4 route get 8.8.8.8 | awk '{ print $5 }')/address").Output()
	return base64.StdEncoding.EncodeToString(bs)
}

func (c *LinuxCollector) vendor() string {
	bs, _ := ioutil.ReadFile("/sys/class/dmi/id/sys_vendor")
	return base64.StdEncoding.EncodeToString(bs)
}

func (c *LinuxCollector) productName() string {
	bs, _ := ioutil.ReadFile("/sys/class/dmi/id/product_name")
	return base64.StdEncoding.EncodeToString(bs)
}

func (c *LinuxCollector) productSerial() string {
	bs, _ := ioutil.ReadFile("/sys/class/dmi/id/product_serial")
	return base64.StdEncoding.EncodeToString(bs)
}

func (c *LinuxCollector) productUUID() string {
	bs, _ := ioutil.ReadFile("/sys/class/dmi/id/product_uuid")
	return base64.StdEncoding.EncodeToString(bs)
}

func (c *LinuxCollector) biosVendor() string {
	bs, _ := ioutil.ReadFile("/sys/class/dmi/id/bios_vendor")
	return base64.StdEncoding.EncodeToString(bs)
}

func (c *LinuxCollector) biosVersion() string {
	bs, _ := ioutil.ReadFile("/sys/class/dmi/id/bios_version")
	return base64.StdEncoding.EncodeToString(bs)
}

func (c *LinuxCollector) biosDate() string {
	bs, _ := ioutil.ReadFile("/sys/class/dmi/id/bios_date")
	return base64.StdEncoding.EncodeToString(bs)
}
