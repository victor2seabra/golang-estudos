package main

import "fmt"

func calculate(valorInicial, taxaJuros float64) (valorFinal float64) {
	valorFinal = valorInicial * (taxaJuros + 1)
	fmt.Printf("O valor final será de: %g", valorFinal)
	return valorFinal
}

func main() {
	calculate(10, 0.9)
}
