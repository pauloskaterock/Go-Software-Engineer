// Em Go, strings são sequências imutáveis de bytes que representam texto. Aqui estão os pontos chave sobre strings em Go:

// Imutáveis: Depois de criada, uma string não pode ser modificada. Qualquer operação que pareça modificar uma string (por exemplo, concatenar ou substituir parte dela) na verdade cria uma nova string.

// Codificação UTF-8: Strings em Go são codificadas em UTF-8, o que significa que podem representar caracteres de qualquer idioma.

// Literais de String: Strings podem ser criadas usando aspas duplas (") ou crases (\).

// Aspas duplas permitem sequências de escape (como \n para nova linha).

// Crases criam strings brutas, onde os caracteres são interpretados literalmente.

// ---------------Exemplos:

package main

import "fmt"

func main() {
	fmt.Println("Olá, Mundo!")
	fmt.Println("Bem-vindo ao seu primeiro projeto em Go!")
}
