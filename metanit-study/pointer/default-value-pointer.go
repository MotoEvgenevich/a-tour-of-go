package main

import "fmt"

func main() {
	p := new(int)
	fmt.Println("value: ", *p) // value: 0 - значение по умолчанию
	*p = 20                    // изменяем значение
	fmt.Println("value: ", *p) // Value: 20
}
