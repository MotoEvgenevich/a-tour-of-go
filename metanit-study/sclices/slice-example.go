package main

import "fmt"

func main() {

	var users []string = []string{"Tom", "Alice", "Kate"}
	fmt.Println(users[2]) //Kate
	users[2] = "Katerine"

	for _, value := range users {
		fmt.Println(value)
	}
}

/* К элементам среза обращение происходит также, как и к элементам массива - по индексу и
также мы можем перебирать все элементы с помощью цикла for:
*/
