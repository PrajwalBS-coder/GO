package temp

import (
	"fmt"
	"errors"
)

func Greet(name string)(string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	msg:=fmt.Sprintf("Hello, %v World!", name)
	return  msg,nil
	// fmt.Println("This is a temporary package.")
}
