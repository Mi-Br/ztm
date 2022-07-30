//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:

//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

//* Create a rectangle structure containing its coordinates
type rectangle struct {
	length, width int
}

//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal

func (r *rectangle) calculateArea() {
	fmt.Println("area:", r.length*r.width)
}
func (r *rectangle) calculatePerimeter() {
	fmt.Println("perimeter:", 2*(r.length+r.width))
}
func (r *rectangle) double() {
	r.length *= 2
	r.width *= 2
}
func main() {
	var r rectangle = rectangle{3, 6}
	r.calculateArea()
	r.calculatePerimeter()
	r.double()
	r.calculateArea()
	r.calculatePerimeter()
}
