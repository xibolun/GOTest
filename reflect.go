package basic

import (
	"reflect"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (p Person) toMap() map[string]interface{} {
	//pV := reflect.ValueOf(p).Elem()
	//
	//for i := 0; i < pV.NumField(); i++ {
	//	vv := pV.Field(i)
	//	kk := pV.Type().Field(i)
	//
	//	mapper[kk.Name] = vv
	//}
	//return mapper

	mapper := make(map[string]interface{})
	vv := reflect.ValueOf(p)
	kk := reflect.TypeOf(p)

	for i := 0; i < kk.NumField(); i++ {
		if len(kk.Field(i).Tag) == 0 {
			mapper[kk.Field(i).Name] = vv.Field(i).Interface()
			continue
		}
		mapper[kk.Field(i).Tag.Get("json")] = vv.Field(i).Interface()
	}
	return mapper
}

func AddPerson() {
	person := Person{
		Name: "zhangsan",
		Age:  10,
	}

	fmt.Println(reflect.TypeOf(person))

	r := reflect.ValueOf(&person).Elem()

	fmt.Println(r)
	fmt.Println(fmt.Sprintf("the field num is : %d", r.NumField()))

	field0 := r.Field(0)
	field1 := r.Field(1)

	fmt.Println(fmt.Sprintf("---------- field 0 is : %s", field0))
	fmt.Println(fmt.Sprintf("---------- field 1 is : %v", field1))

	fmt.Println("----------- start to change value  ---------------")

	s := "pengganyu"
	r.Field(0).SetString(s)
	field0 = r.Field(0)
	fmt.Println(fmt.Sprintf("---------- field 0 is : %s", field0))
	r.Field(1).SetInt(20)
	field0 = r.Field(1)
	fmt.Println(fmt.Sprintf("---------- field 1 is : %v", field1))

}
