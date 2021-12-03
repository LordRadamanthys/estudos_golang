package main

import "math"

// Iniciando com letra maiuscula é publico(visivel dentro e fora do pacote)
// Iniciando com letra minuscula é privado(visivel dentro do pacote)

// Exemplo
// Ponto - sera publico
// ponto ou _Ponto - sera privado

// Ponto representa uma coordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func cateto(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Distancia é responsavel por calcular a distancia linear entre dois ponto
func Distancia(a, b Ponto) float64 {
	cx, cy := cateto(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
