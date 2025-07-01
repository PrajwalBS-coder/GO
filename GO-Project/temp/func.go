package temp

import (
	"fmt"
	"errors"
)

func Greet(name string)(string, error) {
	fmt.Println("Hello, World!", name)
	// fmt.Println("This is a temporary package.")
}
