package basic

import (
	"strconv"
	"testing"
)

func TestResetMap(t *testing.T) {

	type Person struct {
		Name string
		Age  int
	}

	type Class struct {
		Persons map[string]*Person
	}

	var class = Class{
		Persons: make(map[string]*Person, 0),
	}
	for i := 0; i < 10; i++ {
		s := "name" + strconv.Itoa(i)
		class.Persons[s] = &Person{Name: s, Age: i}
	}

	println(len(class.Persons))

	class.Persons = make(map[string]*Person, 0)
	println(len(class.Persons))

}
