package main

import "fmt" // importando librerias

func swap(x string, y string) (string, string) {
	return y, x // retorno multiple
}

func main() {
	a, b := swap("hello", "world") // tipado estático, inferencia de tipos
	fmt.Println(a, b)              // funciones con argumentos variables
	fmt.Println(swap(a, b))
}
