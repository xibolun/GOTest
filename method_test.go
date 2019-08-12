package basic

import "testing"

func TestPolymorphic(t *testing.T) {
	ad := admin{
		user: user{
			name:  "john",
			email: "john@163.com",
		},
		level: "super",
	}

	ad.EchoHello()
}
