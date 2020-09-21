package main

import "fmt"

func main2() {
	var lado float64

	fmt.Println("Dame el valor del un lado del cuadradrado: ")

	fmt.Scanf("%f", &lado)

	fmt.Println("El area del cuadradrado es: ", (lado * lado))

}
