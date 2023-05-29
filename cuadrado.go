package figuras

// Las variables y funciones que comienzan con mayúscula son públicas
type Cuadrado struct {
	Lado float32
}

// mientras que las que comienzan con minúscula son privadas
func (c *Cuadrado) area() float32 {
	return c.Lado * c.Lado
}
func (c *Cuadrado) perimetro() float32 {
	return 4 * c.Lado
}
