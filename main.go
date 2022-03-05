package main

import "fmt"

func main() {
	fmt.Println("Arbol con Espacios: ")
	arbol(7)

}

func arbol(altura int) {
	var i int
	var j int
	n := altura
	for i = 0; i <= n; i++ {
		fmt.Println("")
		// izq
		for j = 1; j <= n-i; j++ {
			fmt.Print(j)
		}
		// med
		for j = 0; j < 2*i-1; j++ {
			fmt.Print(" ")
		}
		// der
		for j = n - i; j > 0; j-- {
			if j != n {
				fmt.Printf("%d", j)
			}
		}
	}

}
