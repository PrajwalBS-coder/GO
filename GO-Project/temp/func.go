package temp

import (
	"fmt"
	"errors"
	"math/rand"
)

func Greet(name string)(string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	msg:=fmt.Sprintf(randomFormat(), name)
	return  msg,nil
}
func randomFormat() string {
    // A slice of message formats.
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }
	return formats[rand.Intn(len(formats))]
}
