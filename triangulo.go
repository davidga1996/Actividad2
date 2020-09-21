package main

import "fmt"

func main3() {
	var base, altura float64

	fmt.Println("Dame la base del triangulo: ")

	fmt.Scanf("%f", &base)

	fmt.Println("Dame la altura del triangulo: ")

	fmt.Scanf("%f", &altura)

	fmt.Println("El area del triangulo es: ", (base * altura / 2))

}
