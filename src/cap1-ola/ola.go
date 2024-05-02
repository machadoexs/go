package main

import "fmt"

func main() {
	bytesnumbers, errors := fmt.Println("Olá, Mundo! Com Go :)")
	fmt.Println(bytesnumbers, errors)

	x := 10
	y := "olá mundo"

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T ", y, y)
}
