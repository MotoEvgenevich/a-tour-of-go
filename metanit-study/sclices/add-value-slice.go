package main

import "fmt"

func main() {
	users := []string{"Tom", "Alice", "Kate"}
	fmt.Println(users)
	users = append(users, "Bob")

	for _, value := range users {
		fmt.Println(value)
	}
}

/* Для добавления в срез применяется встроенная функция append(slice, value).
Первый параметр функции - срез, в который надо добавить, а второй параметр - значение,
которое нужно добавить. Результатом функции является увеличенный срез. */
