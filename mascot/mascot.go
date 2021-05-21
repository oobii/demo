package mascot

import "fmt"

// In Go, a function whose name starts with a capital letter can be called by a function not in the same package
// BestMascot return best mascort name
func BestMascot() string {
	return "Go Gopher"
}

func Greeting(greeting string) string {
	// In Go, the := operator is a shortcut for declaring and initializing a variable in one line
	message := fmt.Sprintf("%v %v!",greeting,BestMascot())
	return message
}