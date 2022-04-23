package main

import "fmt"

func main() {
	users := []string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	fmt.Println(users)
	// удаляем 4 элемент
	var n = 3
	users = append(users[:n], users[n+1:]...)
	fmt.Println(users)

}

// Что делать, если необходимо удалить какой-то определенный элемент?

// В этом случае мы можем комбинировать функцию append и оператор среза
