package basic

import (
	"fmt"
	"regexp"
	"testing"
)

func TestNumber(t *testing.T) {
	fmt.Println(Match(`^[0-9]+$`, "032413"))
}

func TestCheckSyntax(t *testing.T) {
	reg := `([^\s\w])(\s*\1)+`
	if _, err := regexp.Compile(reg); err != nil {
		t.Error(err.Error())
	}

}

func TestPhone(t *testing.T) {
	reg := `^(((\\+\\d{2}-)?0\\d{2,3}-\\d{7,8})|((\\+\\d{2}-)?(\\d{2,3}-)?([1][3,4,5,7,8][0-9]\\d{8})))$`
	fmt.Println(Match(reg, "133"))
	fmt.Println(Match(reg, "1803710573"))
	fmt.Println(Match(reg, "18037105737"))
	fmt.Println(Match(reg, "0571-83586104"))
	fmt.Println(Match(reg, "0571-8358610"))
	fmt.Println(Match(reg, "057183586104"))
	fmt.Println(Match(reg, "83586104"))

}

func TestString(t *testing.T) {
	//reg := `^/api/cloudboot/v1/devices/([0-9A-Za-z_]+|(?!environment))/collections$`
	reg := `^/api/cloudboot/v1/devices/[^env][0-9A-Za-z_]+/collections$`
	//reg2 := `(?!env\w)+`
	fmt.Println(Match(reg, "/api/cloudboot/v1/devices/3241234env/collections"))
	fmt.Println(Match(reg, "/api/cloudboot/v1/devices/1111111/collections"))
	fmt.Println(Match(reg, "/api/cloudboot/v1/devices/env12333/collections"))
	//fmt.Println(Match(reg2, "/api/cloudboot/v1/devices/123/collections"))
}

func Match(reg, str string) bool {
	match, err := regexp.MatchString(reg, str)
	return match == true && err != nil
}

func TestIP(t *testing.T) {
	reg := `^((2[0-4]\d|25[0-5]|[01]?\d\d?)\.){3}(2[0-4]\d|25[0-5]|[01]?\d\d?)$`
	//reg := `\\d{0,3}\\.\\d{0,3}\\.\\d{0,3}\\.\\d{0,3}`
	//reg := `（\d{0,3}\.){3}\d{0,3}`
	//reg := "^((2[0-4]\\d|25[0-5]|[01]?\\d\\d?)\\.){3}(2[0-4]\\d|25[0-5]|[01]?\\d\\d?)$"

	ipReg := regexp.MustCompile(reg)
	fmt.Println(ipReg.MatchString("10.0.0.259"))
}
