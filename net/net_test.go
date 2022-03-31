package main

import (
	"fmt"
	"net"
)

func main() {
	nics, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	for i := range nics {
		fmt.Printf("%v\n", nics[i])
	}

}
