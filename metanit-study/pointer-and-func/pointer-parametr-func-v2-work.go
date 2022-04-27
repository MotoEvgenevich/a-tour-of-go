package main

import "fmt"

func changeValue(x *int) {
	*x = (*x) * (*x)

}

func main() {
	d := 5
	fmt.Println("d Before:", d) //5
	changeValue(&d)             // изменяем значение
	fmt.Println("d After:", d)  // 25 -значение изменилось!
}
