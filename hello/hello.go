package hello

import (
	"fmt"

	"github.com/abdullahtarek/component_dependency_go_2/greetings"
)

func Hello() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
