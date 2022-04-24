package main

import "fmt"

func main() {

	var people = map[string]int{
		"Tom":   1,
		"Bob":   2,
		"Sam":   4,
		"Alice": 8,
	}
	fmt.Println(people["Alice"])
	fmt.Println(people["Bob"])
	fmt.Println(people["Sam"])
	people["Bob"] = 32
	fmt.Println(people["Bob"])
	fmt.Println("if key = true; display value = 1")
	ifSample()
	fmt.Println("Для перебора элементов используется цикл for:")
	forSample()
	fmt.Println("make emty map and for range value:")
	makeSample()
	fmt.Printf("добавляем элемент в таблицу people[Kate] = 128\n")
	addElement()
	fmt.Printf("delete \"Bob\" из мапы: выводим оставшихся\n")
	deleteSample()
}

/* Как и в массиве или в срезе элементы помещаютя в фигурные скобки. Каждый элемент представляет пару ключ -значение.
Вначале идет ключ и через двоеточие значение. Определение элемента завершается запятой.

Для обращения к элементам нужно применять ключи:
*/

func ifSample() {
	var people = map[string]int{
		"Tom":   1,
		"Bob":   2,
		"Sam":   4,
		"Alice": 8,
	}
	if val, ok := people["Tom"]; ok {
		fmt.Println(val)
	}
}

/* Если значение по заданному ключу имеется в отображении, то переменная ok будет равна true,
а переменная val будет хранить полученное значение. Если переменная ok равна false,
то значения по ключу в отображении нет. */

func forSample() {
	var people = map[string]int{
		"Tom":   1,
		"Bob":   2,
		"Sam":   4,
		"Alice": 8,
	}
	for key, value := range people {
		fmt.Println(key, value)
	}
}

func makeSample() {
	people := make(map[string]int)
	people["Koka"] = 32

	for key, value := range people {
		fmt.Println(key, value)
	}
}

func addElement() {
	var people = map[string]int{"Tom": 1, "Bob": 2}
	people["Kate"] = 128
	fmt.Println(people) // map[Tom:1  Bob:2  Kate:128]
}

func deleteSample() {
	var people = map[string]int{"Tom": 1, "Bob": 2, "Sam": 8}
	delete(people, "Bob")
	fmt.Println(people)
}
