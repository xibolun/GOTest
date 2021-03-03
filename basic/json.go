package basic

import (
	json2 "encoding/json"
	"fmt"
	"strings"
	"reflect"
	"time"
)

func JsonArray2Struct(json, name string) {

	var arr []map[string]interface{}

	json2.Unmarshal([]byte(json), &arr)

	convertStruct(arr[0], name)

}

func convertStruct(obj map[string]interface{}, name string) {
	fmt.Printf("%s struct{ \n", ToStructKey(name))
	for k, v := range obj {
		var structType interface{}
		rt := reflect.TypeOf(v)
		switch rt.Kind() {
		case reflect.String:
			structType = "string"
		case reflect.Bool:
			structType = "bool"
		case reflect.Int:
			structType = "int"
		case reflect.Map:
			go JsonObj2Struct(ToJsonString(v), ToStructKey(k))
			structType = "[]" + ToStructKey(k)
		case reflect.Slice:
			go JsonArray2Struct(ToJsonString(v), ToStructKey(k))
			structType = "[]" + ToStructKey(k)
		case reflect.Array:
			structType = "[]string"
		default:
			structType = "string"
		}

		fmt.Printf("%s\t%s\t%s\n", ToStructKey(k), structType, fmt.Sprintf("`json:\"%s\"`", k))
	}
	fmt.Println("}")

	time.Sleep(500 * time.Microsecond)
}

func JsonObj2Struct(json, name string) {
	obj := make(map[string]interface{})

	json2.Unmarshal([]byte(json), &obj)

	convertStruct(obj, name)

}

func ToStructKey(str string) (new string) {
	return IDUpper(FirstUpper(str))
}

func FirstUpper(str string) (new string) {

	firstVal := string(str[0])

	firstUpper := strings.ToUpper(firstVal)

	return strings.Replace(str, firstVal, firstUpper, 1)
}

func IDUpper(str string) (new string) {
	str = strings.Replace(str, "Id", "ID", 1)
	return strings.Replace(str, "id", "ID", 1)
}

/**
value to json string
*/
func ToJsonString(v interface{}) string {
	bytes, err := json2.Marshal(v)
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(bytes)
}

// TODO xorm struct to json
//func XormToJSON(xorm interface{}) string {
//
//}
