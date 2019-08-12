package basic

import "fmt"

type user struct {
	name  string
	email string
}

func (u *user) EchoHello() {
	fmt.Println("hello ")
}

type admin struct {
	user
	level string
}


