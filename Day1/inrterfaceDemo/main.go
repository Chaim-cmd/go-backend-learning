package main

import (
	"fmt"
	"math"
)

// 定义接口
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width, Height float64
}

// circle 实现Shape
func (c Circle) Area() float64      { return math.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }

// Rectangle 实现Shape
func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width * r.Height) }

// 函数接收接口，不关心具体类型
func printShapeInfo(s Shape) {
	fmt.Printf("面积： %.2f,周长：%.2f\n", s.Area(), s.Perimeter())
}

func main() {
	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 4, Height: 6},
	}

	for _, v := range shapes {
		//类型断言
		switch v := v.(type) {
		case Circle:
			fmt.Println("圆形的半径：", v.Radius)
		case Rectangle:
			fmt.Println("矩形：", v.Width, v.Height)
		}
		printShapeInfo(v)

	}
}
