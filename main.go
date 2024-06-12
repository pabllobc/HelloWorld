package main

import (
	"fmt"
	"os"
)

func main() {
	var n1 float32
	var n2 float32

	n1 = 100
	n2 = 220
	soma := n1 + n2

	fmt.Fprintln(os.Stdout, []any{"Soma = %f", soma}...)

	fmt.Println("Hello World!!!")
}
