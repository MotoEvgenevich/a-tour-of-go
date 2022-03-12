package main 
import "fmt"

func main() {
	fmt.Println("Enter 3 numbers :")
	var first int
	fmt.Scanln(&first)
	var second int
	fmt.Scanln(&second)
	var third int
	fmt.Scanln(&third)
	/*check the boolean condition using if statement
	провреряем логическое сравнение с помощью оператора if*/
	if first >= second && first >= third {
		fmt.Println(first, "is Largest amoung three numbers.")
		/* if condition is true then print the following
		Если условие истино выведи следующее */
	}
	if second >= first && second >= third {
		fmt.Println(second, "is Largest amoung three numbers.")
	}
	if third >= first && third >= second {
		fmt.Println(third, "is Largest amoung three numbers.")
	}
	if first = second && first = third {
		fmt.Println("this number's is equal")
	}
}