package main

import "fmt"

func main() {
	//source massive
	initialUsers := [8]string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	users1 := initialUsers[2:6] // 3 по 7
	users2 := initialUsers[:4]  // 1 по 3
	users3 := initialUsers[3:]  // с 4 до конца
	users4 := initialUsers[:]   // просто интересно пересоздать слайс из массива

	fmt.Println(users1) // ["Kate", "Sam", "Tom", "Paul"]
	fmt.Println(users2) // ["Bob", "Alice", "Kate", "Sam"]
	fmt.Println(users3) // ["Sam", "Tom", "Paul", "Mike", "Robert"]
	fmt.Println(users4) // ["Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"]
}

/* Оператор среза s[i:j] создает из последовательности s новый срез,
который содержит элементы последовательности s с i по j-1.
При этом должно соблюдаться условие 0 <= i <= j <= cap(s). В качестве исходной последовательности,
из которой берутся элементы, может использоваться массив, указатель на массив или другой срез.
В итоге в полученном срезе будет j-i элементов.

Если значение i не указано, то применяется по умолчанию значение 0.
Если значение j не указано, то вместо него используется длина исходной последовательности s.
*/
