package main

import (
	"fmt"
	"math"
)

type Cilindro struct {
	// atributos
	Radio  float64
	Altura float64
}

// metodo(s)
func (c Cilindro) Volumen() float64 {
	base := math.Pi * math.Pow(c.Radio, 2)
	return base * c.Altura
}

func main() {
	cuerpo := Cilindro{Radio: 5, Altura: 10}
	fmt.Println(cuerpo.Volumen())
}
