package main

import "fmt"

func main() {
	// Booleanos em Go representam valores verdadeiro (true) ou falso (false).
	var verdadeiro bool = true // declaração explícita
	var falso = false          // declaração implícita (Go infere o tipo)

	fmt.Println("Valor de verdadeiro:", verdadeiro)
	fmt.Println("Valor de falso:", falso)

	// Operadores relacionais retornam booleanos
	a := 10
	b := 20
	resultado := a < b // resultado será true, pois 10 é menor que 20
	fmt.Println("a < b:", resultado)

	// Operadores lógicos: && (E), || (OU), ! (NÃO)
	x := true
	y := false
	fmt.Println("x && y:", x && y) // false (ambos precisam ser true)
	fmt.Println("x || y:", x || y) // true (basta um ser true)
	fmt.Println("!x:", !x)         // false (negação)

	// Zero value: o valor padrão de um booleano é false
	var padrao bool
	fmt.Println("Valor padrão de bool:", padrao)
}
