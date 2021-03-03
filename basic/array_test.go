package basic

import (
	"fmt"
	"testing"
)

func TestArrayPoint(t *testing.T) {
	sp := [3]*string{new(string), new(string), new(string)}
	*sp[0] = "red"
	*sp[1] = "blue"
	*sp[2] = "green"

	for i, _ := range sp {
		fmt.Printf("point value %v,real value %s\n", sp[i], *sp[i])
	}
}

