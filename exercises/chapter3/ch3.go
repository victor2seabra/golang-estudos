package main

import "fmt"

func arrays_for_test() {
	var x [5]int64                    // cria uma array no formato {0, 0, 0, 0, 0}
	var y [5]string                   // cria uma array no formato {"", "", "", "", ""}
	d := [2]string{"olá", "tudo bom"} // cria uma array de len 2 e cap 2

	var b []string           // cria um slice no formato {}. valida true quando comparado a nil
	var c = []string{}       // cria uma array no formato {}. valida false quando comparado a nil
	var z = make([]int64, 5) // cria um slice no formato {0, 0, 0, 0, 0}
	a := make([]string, 5)   // cria um slice no formato {"", "", "", "", ""}

	fmt.Printf("%v", x)
	fmt.Print(y)
	fmt.Print(z)
	fmt.Print(a)
	fmt.Print(b)
	fmt.Print(d)
	fmt.Print(c)
}

func main() {
	arrays_for_test()
}
