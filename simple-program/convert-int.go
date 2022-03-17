package main

import (
	"fmt"
	"strconv"
)

const bin = "11111111"
const hex = "1A"
const oct = "12"
const dec = "10"
const floatNum = "16.123557"

func main() {
	// Конвертация двоичного значения в шеснадцатиричное
	v, _ := ConvertInt(bin, 2, 16)
	fmt.Printf("Двоичное значение %s конвертировано в шеснадцатиричное: %s\n", bin, v)

	// Конвертация шеснадцатиричного значения в десятичное
	v, _ = ConvertInt(hex, 16, 10)
	fmt.Printf("Шеснадцатиричное значение %s конвертировано в десятичное: %s\n", hex, v)

	// Конвертация восьмеричного значения в шеснадцатиричное
	v, _ = ConvertInt(oct, 8, 16)

	// Конвертация десятиричного значения в восьмеричное
	v, _ = ConvertInt(dec, 10, 8)
	fmt.Printf("Десятиричное значение %s конвертировано в восьмеричное: %s\n", dec, v)

	// а теперь задание из учебника на наглядном примере:
	// Вопрос: каково самое большое число из 8 цифр в двоичной системе
	// "11111111"нужно перевести в десятичную систему

	v, _ = ConvertInt(bin, 2, 10)
	fmt.Printf("Двоичное значение %s конвертировано в десятичное: %s\n", bin, v)

}

// ConvertInt конвертирует значение из одной системы счисления
// в другую, которая указана в toBase

func ConvertInt(val string, base, toBase int) (string, error) {
	i, err := strconv.ParseInt(val, base, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, toBase), nil
}
