package basic

import (
	"fmt"
	"reflect"
	"testing"
)

// https://juejin.im/post/5a75a4fb5188257a82110544

func TestAddPerson(t *testing.T) {
	//AddPerson()
}

func TestPersonReflect(t *testing.T) {

	per := Person{
		Name: "zhangsan",
		Age:  20,
	}

	perP := &per

	perV := reflect.ValueOf(perP)

	a := perV.Elem()

	fmt.Printf("poing value elem type is :%s\n", a.Type().Kind())

	for i := 0; i < a.Type().NumField(); i++ {
		fmt.Println(a.Type().Field(i).Name)
	}

	fmt.Println(a.FieldByName("Name"))

	fmt.Println(perV.Type())
	fmt.Println(perV.Type().Kind())

	fmt.Println(perV.Elem())

}
