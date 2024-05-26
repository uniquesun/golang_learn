package main

import "fmt"

// Shaper 图形接口
type Shaper interface {
	Area() float32
}

// Square 正方形
type Square struct {
	side float32
}

// Area 类型实现了图形接口
func (s *Square) Area() float32 {
	return s.side * s.side
}

// Rectangle 长方形
type Rectangle struct {
	length, width float32
}

func (r *Rectangle) Area() float32 {
	return r.length * r.width
}

func main() {
	sq1 := new(Square)
	sq1.side = 5

	// 接口类型
	var areaIntf Shaper
	fmt.Println(areaIntf) // nil
	areaIntf = sq1

	// 接口断言
	if _, ok := areaIntf.(*Square); ok {
		fmt.Printf("是Square")
	}

	if _, ok := areaIntf.(*Rectangle); ok {
		fmt.Printf("是 Rectangle")
	}

	if _, ok := areaIntf.(Shaper); ok {
		fmt.Println("是 Shaper")
	}

	// 类型断言
	switch t := areaIntf.(type) {
	case *Square:
		fmt.Printf("Type Square %T with value %v\n", t, t)
	case *Rectangle:
		fmt.Printf("Type Circle %T with value %v\n", t, t)
	case nil:
		fmt.Printf("nil value: nothing to check?\n")
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}

	fmt.Printf("The square has area: %f\n", areaIntf.Area())

	// 多态
	//r := &Rectangle{5, 3}
	//q := &Square{5}
	//
	//shapes := []Shaper{r, q}
	//fmt.Println("Looping through shapes for area ...")
	//for n, _ := range shapes {
	//	fmt.Println("Shape details: ", shapes[n])
	//	fmt.Println("Area of this shape is: ", shapes[n].Area()) // Area方法走不通的方法
	//}
}
