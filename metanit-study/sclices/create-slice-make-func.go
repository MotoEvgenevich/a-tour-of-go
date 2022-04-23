package main

import "fmt"

var users []string = make([]string, 3)

func main() {

	users[0] = "Tom"
	users[1] = "Alice"
	users[2] = "Bob"
	fmt.Println(users[0])
	fmt.Println(users[1])
	fmt.Println(users[2])

	for _, value := range users {
		fmt.Println(value)
	}

}

// С помощью функции make() можно создать срез из нескольких элементов, которые будут иметь значения по умолчанию ""
