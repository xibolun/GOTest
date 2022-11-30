package basic

import (
	"encoding/binary"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestConvert(t *testing.T) {
	fmt.Println(byteArrToInt(intToByteArr(100000000000000)))
	fmt.Println(byteArrToBool(boolToByteArr(true)))

	fmt.Println(strConvert("1", reflect.Int.String()))
	fmt.Println(strConvert("1", reflect.Uint.String()))
	fmt.Println(strConvert("2018-11-28 01:21:06 +0800 CST", "time"))
}

func byteArrToInt(bytes []byte) int {
	return int(binary.BigEndian.Uint64(bytes))

}

func byteArrToBool(bytes []byte) bool {
	b, _ := strconv.ParseBool(string(bytes))
	return b
}

func boolToByteArr(bool bool) []byte {
	return []byte(strconv.FormatBool(bool))
}

func intToByteArr(int int) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(int))
	return buf
}

func strConvert(str, typ string) (v interface{}, err error) {
	switch typ {
	case reflect.Int.String():
		v, err = strconv.ParseInt(str, 10, 0)
	case reflect.Int8.String():
		v, err = strconv.ParseInt(str, 10, 8)
	case reflect.Int16.String():
		v, err = strconv.ParseInt(str, 10, 16)
	case reflect.Int32.String():
		v, err = strconv.ParseInt(str, 10, 32)
	case reflect.Int64.String():
		v, err = strconv.ParseInt(str, 10, 64)
	case reflect.Uint.String():
		v, err = strconv.ParseUint(str, 10, 0)
	case reflect.Uint8.String():
		v, err = strconv.ParseUint(str, 10, 8)
	case reflect.Uint16.String():
		v, err = strconv.ParseUint(str, 10, 16)
	case reflect.Uint32.String():
		v, err = strconv.ParseUint(str, 10, 32)
	case reflect.Uint64.String():
		v, err = strconv.ParseUint(str, 10, 64)
	case reflect.Bool.String():
		v, err = strconv.ParseBool(str)
	case reflect.Float32.String():
		v, err = strconv.ParseFloat(str, 32)
	case reflect.Float64.String():
		v, err = strconv.ParseFloat(str, 64)
	case "reader":
		v = strings.NewReader(str)
	case "time":
		if strings.Contains(str, "CST") {
			v, err = time.Parse("2006-01-02 15:04:05 +0800 CST", str)
		}
	default:
		v = str
	}
	fmt.Println(reflect.TypeOf(v))
	return
}

func strToInt(str string) int {
	v, _ := strconv.Atoi(str)
	return v
}

func strToBool(str string) bool {
	v, _ := strconv.ParseBool(str)
	return v
}

func TestFloat64ToInt(t *testing.T) {
	ast := assert.New(t)

	vv := map[float64]int{
		3.2:      3,
		3.200001: 3,
		3.500001: 4,
		0:        0,
		-3.2:     -3,
		-3.9:     -4,
	}

	for k, v := range vv {
		ast.Equal(v, int(math.Round(k)))
	}
}
