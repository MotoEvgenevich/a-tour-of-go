package main

import "fmt"

func main() {
	fmt.Print("Введите длинну в футах:")
	var foot float32
	const m float32 = 0.3048
	fmt.Scan(&foot)
	foot *= m
	fmt.Print(foot, ":Метра(ов)")

}
