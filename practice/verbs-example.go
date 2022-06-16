package main

import (
	"fmt"
)

var integer int64 = 32500
var floatNum float64 = 22000.456

func main() {

	// Обычный способ вывода десятичного числа
	fmt.Printf("%d \n", integer)

	// Всегда показывает знак
	fmt.Printf("%+d \n", integer)

	// Вывод с другим основанием x -16, o-8, b -2, d - 10
	fmt.Printf("%x \n", integer)
	fmt.Printf("%#x \n", integer)

	// Отступ перед нулями
	fmt.Printf("%010d \n", integer)

	// Оставляет отступ с пробелами
	fmt.Printf("% 10d \n", integer)

	// Отступ с правой стороны
	fmt.Printf("% -10d \n", integer)

	// Вывод вещественного значения
	// с плавающей запятой
	fmt.Printf("%f \n", floatNum)

	// Вещественное число
	// с ограниченной точностью = 5 (после запятой)
	fmt.Printf("%.5f \n", floatNum)

	// Вещественное число
	// в научной заметке
	fmt.Printf("%e \n", floatNum)

	// Вещественное число
	// %e для крупной экспоненты
	// или %f в противном случае
	fmt.Printf("%g \n", floatNum)

}
