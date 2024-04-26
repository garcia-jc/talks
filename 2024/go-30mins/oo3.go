package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	vaso := Vaso{Cilindro: Cilindro{Radio: 5, Altura: 10}}
	vaso.agregarLiquido(700)
	err := vaso.agregarLiquido(100)
	fmt.Println(vaso, err)
}

type Vaso struct {
	Cilindro         // herencia?
	volumenContenido float64
}

func (v *Vaso) agregarLiquido(volumen float64) error { //encapsulación
	if v.volumenContenido+volumen > v.Volumen() {
		return errors.New("el vaso se derramará!")
	}
	v.volumenContenido += volumen
	return nil
}

func (v Vaso) String() string { // polimorfismo
	return fmt.Sprintf("vaso(tamaño: %v mL, usado: %v mL", v.Volumen(), v.volumenContenido)
}

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
