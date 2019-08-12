package basic

import (
	"reflect"
	"fmt"
)

type Person struct {
	Name string
	Age  int
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
