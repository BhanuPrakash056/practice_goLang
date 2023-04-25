//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type coordinates struct {
	x, y int
}
type rectangle struct {
	a coordinates
	b coordinates
}

func width(rect rectangle) int {
	return (rect.b.x - rect.a.x)
}

func length(rect rectangle) int {
	return (rect.a.y - rect.b.y)
}
func area(rect rectangle) int {
	return length(rect) * width(rect)
}
func perimeter(rect rectangle) int {
	return 2 * (length(rect) + width(rect))
}
func printInfo(rect rectangle) {
	fmt.Println("Area is ", area(rect))
	fmt.Println("Parameter is ", perimeter(rect))
}
func main() {
	rect := rectangle{a: coordinates{0, 7}, b: coordinates{10, 0}}
	printInfo(rect)
	rect.a.y *= 2
	rect.b.x *= 2
	printInfo(rect)
}
