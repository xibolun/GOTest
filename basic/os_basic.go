package basic

import (
	"fmt"
	u "os/user"
)

func OSUser() {
	// user
	u1, _ := u.Lookup("admin")
	fmt.Printf("user.Lookup info : %s", ToJsonString(u1))

	// cureent user
	current, _ := u.Current()
	fmt.Printf("current user : %s", ToJsonString(current))
}
