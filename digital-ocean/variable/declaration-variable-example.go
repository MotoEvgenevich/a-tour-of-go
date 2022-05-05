package main

import "fmt"

var f bool
var G string
var G_1 int // символы подчеркивания разрешены
// j, k, l := "shark", 2.05, 25 				!!!!!!!!!ЗДЕСЬ НЕ МОГУ ТАК ДЕКЛАРИРОВАТЬ ТОЛЬКО В ТЕЛЕ ФУНКЦИИ!!!!!!
func main() {
	var a int
	var b string
	var c float64
	var d bool
	var e int64
	j, k, l := "shark", 2.05, 25

	fmt.Printf("var a %T = %+v\n", a, a)
	fmt.Printf("var b %T = %q\n", b, b)
	fmt.Printf("var c %T = %+v\n", c, c)
	fmt.Printf("var d %T = %+v\n\n\n", d, d) // опечатка на дижитал оушене чем больше \n тем больше отступов новых строк
	fmt.Printf("var e %T = %+v\n", e, e)
	fmt.Printf("var f %T = %v\n", f, f)
	fmt.Printf("var G %T = %q\n", G, G)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)
}
