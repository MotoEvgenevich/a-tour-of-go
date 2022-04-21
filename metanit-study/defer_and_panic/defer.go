package main

import "fmt"

func main() {
	defer finish()
	fmt.Println("Program has been started")
	fmt.Println("Program is working")
}

func finish() {
	fmt.Println("Pogram has been finished")
}

/* Оператор defer позволяет выполнить определенную функцию в конце программы, при этом не важно,
где в реальности вызывается эта функция. Здесь функция finish вызывается с оператором defer,
поэтому данная функция в реальности будет вызываться в самом конце выполнения программы,
несмотря на то, что ее вызов определен в начале функции main. */
