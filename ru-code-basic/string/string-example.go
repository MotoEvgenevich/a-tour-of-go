package main

import (
	"fmt"
	"strings"
)

func main() {

}

func Greetings(name string) string {
	name = strings.Trim(name, " ")
	name = strings.ToLower(name)
	name = strings.Title(name)

	return fmt.Sprintf("Привет, %s!", name)
}
