package screens

import "fmt"

func Prompt(message string) string {
	var user_input string
	fmt.Print(message)
	fmt.Scanln(&user_input)
	return user_input
}
