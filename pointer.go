package basic

import (
	"math"
)
//
//func main() {
//	v := Vertex{3, 4}
//	v.Scale(2)
//	//// go语言不能传引用 ，只能传引用所带的指针
//	ScaleFunc(&v, 10) /////  此处必须使用v的指针，因为v是一个引用 ，而非真实的Vertex
//
//	i := 42
//	fmt.Println(&i)
//	fmt.Println(i)
//
//	fmt.Println(v)
//}

/**
 * 定义构造方法
 */
type Vertex struct {
	X, Y float64
}

/**
 * Vertex构造方法的Scale属性
 * @param  {[type]} v *Vertex)      Scale (f float64 [description]
 * @return {[type]}   [description]
 */
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/**
 * 构造函数做为参数进行传递
 * @param {[type]} v *Vertex [description]
 * @param {[type]} f float64 [description]
 */
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
