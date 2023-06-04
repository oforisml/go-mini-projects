package main

import "fmt"



type area interface{
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	base float64
	height float64
}


func main (){
	s := square{sideLength: 0.5}
	t := triangle{base: 0.5, height: 2}
	printArea(s)
	printArea(t)

}



func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(a area) {
	fmt.Println(a.getArea())
}