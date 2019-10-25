package basic

import (
	"fmt"
	"sync"
	"testing"

	"github.com/mitchellh/mapstructure"
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

	userAge.Add("zhangsan", 10)
	fmt.Println(userAge.Get("zhangsan"))
}

func TestMapNil(t *testing.T) {
	var mapA map[string]string

	mapA["name"] = "zhangsan"

	t.Logf("map value %v", mapA)
}

func TestMapOper(t *testing.T) {
	mapA := make(map[string]string)
	mapA["name"] = "zhangsan"
	mapA["age"] = "10"
	mapA["address"] = "浙江杭州"
	mapA["company"] = "中科院"

	// map len
	t.Logf(" mapA length is %d", len(mapA))

	// get value
	name := mapA["name"]
	t.Logf("name is %s", name)

	// is exist
	age, ok := mapA["age"]
	t.Logf("age value is exist? %t, value is %s", ok, age)

	// delete one
	delete(mapA, "company")
	t.Logf("mapA %v", mapA)

	// loop
	for k, v := range mapA {
		t.Logf(" k: %s, v : %s", k, v)
	}
}

func TestMapParallel(t *testing.T) {
	t.Parallel()

	mapA := make(map[int]int)

	t.Run("a", func(t *testing.T) {
		for i := 0; i < 100; i++ {
			mapA[i] = i
		}
		t.Logf("mapA %v", mapA)
	})
}

func TestMapClosure1(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	for k, v := range m {
		go func(kk *string, vv *int) {
			fmt.Printf("%v,%v\n", *kk, *vv)
		}(&k, &v)
	}
}

func TestMapClosure2(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	for k, v := range m {
		fmt.Printf("origin: %v,%v, %s,%d\n", &k, &v, k, v)
		go func(kk *string, vv *int) {
			fmt.Printf("closure: %v,%v, %s,%d\n", kk, vv, *kk, *vv)
		}(&k, &v)
	}
}
