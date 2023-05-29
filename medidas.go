package figuras

import "fmt"

type Figuras interface {
	area() float32
	perimetro() float32
}

func Medidas(f Figuras) {
	fmt.Println("Area:", f.area())
	fmt.Println("Perimetro:", f.perimetro())

}
