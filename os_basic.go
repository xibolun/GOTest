package basic

import (
	"fmt"
	"os/user"
)

func OSUser() {
	// user
	u, _ := user.Lookup("admin")
	fmt.Printf("user.Lookup info : %s", ToJsonString(u))

	// cureent user
	current, _ := user.Current()
	fmt.Printf("current user : %s", ToJsonString(current))
}


