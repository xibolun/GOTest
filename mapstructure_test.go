package basic

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"sync"
	"testing"
)

func Test_mapstructure(t *testing.T) {
	person := map[string]interface{}{
		"Name": "chenjialin",
		"Age":  36,
	}

	var per Person
	mapstructure.Decode(person, &per)
	fmt.Println(ToJsonString(per))
}

func Test_ConcurrentMap(t *testing.T) {
	userAge := UserAges{
		map[string]int{},
		sync.Mutex{},
	}

	userAge.Add("zhangsan",10)
	fmt.Println(userAge.Get("zhangsan"))
}
