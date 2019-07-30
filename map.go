package basic

import (
	"fmt"
	"sync"
)

//func main() {
//	mapFunc()
//}

func mapFunc() {
	var m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.1, -74.2,
	}
	fmt.Println(m)

	// 定义一个mpa对象[string]定义key的类型  后面接value的类型
	// map对象只能由make进行创建，不能new
	var _mapObj = make(map[string]int)
	_mapObj["hello"] = 42
	_mapObj["world"] = 41

	// 利用双赋值来检测某个键的value是否存在
	_value, ok := _mapObj["hello"]
	fmt.Println(_mapObj["hello"])
	fmt.Println(_value, ok) // 42,true

	fmt.Println(_mapObj)

	// 删除map对象当中的一个元素，此函数无返回值
	delete(_mapObj, "world")

	fmt.Println(_mapObj)
}

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}
