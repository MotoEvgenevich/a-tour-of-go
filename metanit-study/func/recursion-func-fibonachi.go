package main

import "fmt"

func fibbonachi(n uint) uint {

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibbonachi(n-1) + fibbonachi(n-2)
}
func main() {

	fmt.Println(fibbonachi(4)) // 3
	fmt.Println(fibbonachi(5)) // 5
	fmt.Println(fibbonachi(6)) // 8
}

/* Другим распространенным показательным примером рекурсивной функции служит функция,
вычисляющая числа Фибоначчи. n-й член последовательности Фибоначчи определяется по формуле:
f(n)=f(n-1) + f(n-2), причем f(0)=0, а f(1)=1.
*/
