package teste

import "fmt"

// definindo explicitamente variáveis
// global scope
const x int64 = 10 // uma constante nao pode ter seu valor alterado

// somente funções e outros objetos iniciados com letra maiscula podem ser importados
func Teste() {
	// definindo variaveis com tipagem explicita dentro de funções
	// local scope
	var y float64 = 5.5

	// definindo variável com tipagem implicita
	// possivel somente dentro de funções
	z := 5

	fmt.Printf("O valor de x é %d.\nO valor de y é %.2f.\nO valor de z é %d\n", x, y, z)

	/*
		num condicional podemos criar uma variavel e já usá-la em seguida
		note que para somar um int e um float foi necessário a conversão
		por conta da forte tipagem da linguagem Go
	*/
	if a := z + int(y); a != 0 {
		fmt.Println("Not zero")
	} else if b := false; !b {
		fmt.Print(b)
	} else {
		fmt.Println("Zero")
	}

	// estrutura de um for loop
	// inicialização; condição; após a condição validar
	for b := y; b > float64(z); b++ {
		fmt.Printf("%v", b)
	}
}
