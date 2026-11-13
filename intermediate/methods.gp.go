package intermediate

import "fmt"

type Rectangle struct {
	length float64
	width  float64
}

func Methods() {
	rect := Rectangle{10, 5}

	area := rect.Area()
	fmt.Println("Area:", area)
}

func (r Rectangle) Area() float64 {
	return r.width * r.length
}
