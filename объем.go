package main

import "fmt"

func main() {
	
	a := 5
	b := 2
	c := 8
	
	V := a * b * c
	
	S := 2 * (a*b + b*c + a*c)
	
	fmt.Println("Объем ", V)
	fmt.Println("Площадь ", S)
}
