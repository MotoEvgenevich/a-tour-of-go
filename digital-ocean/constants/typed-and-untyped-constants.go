package main

import "fmt"

const (
	year     = 365
	leapYear = int32(366)
)

func main() {
	hours := 24
	minutes := int32(60)
	fmt.Println(hours * year)
	fmt.Println(minutes * year)
	fmt.Println(minutes * leapYear)
	//fmt.Println(hours * leapYear) 			//!!!!!!!! Если расскоментить ошибка invalid operation: hours * leapYear (mismatched types int and int32) НЕСОВМЕСТИМЫЕ ТИПЫ ДАННЫХ!!!!!!!!!
}
