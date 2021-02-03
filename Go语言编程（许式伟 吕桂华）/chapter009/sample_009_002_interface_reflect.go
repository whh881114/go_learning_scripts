package main

import "fmt"

type Shaper interface{ Area() float32 }

type Square struct{ side float32 }

func (sq *Square) Area() float32 { return sq.side * sq.side }

type Rectangle struct{ length, width float32 }

func (r Rectangle) Area() float32 { return r.length * r.width }

func main() {
	r := Rectangle{5, 3}
	q := &Square{5}
	// shapes := []Shaper{Shaper(r), Shaper(q)}

	// 更简洁的写法：
	shapes := []Shaper{r, q} // Shaper接口有Area()统一屏蔽了Rectangle和Square相同的函数名。

	// 遍历shapes
	for n, _ := range shapes {
		fmt.Println("形状参数：", shapes[n])
		fmt.Println("形状面积是：", shapes[n].Area())
	}
}
