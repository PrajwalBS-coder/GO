package temp

import (
	"fmt"
	"errors"
	"math/rand"
)

// func Greet(name [] string)([]string, error) {
// 	var a []string
// 	for _, n := range name{
// 		if n ==''{
// 			return nil errors.New("empty name")
// 		}
		
// 		msg:=fmt.Sprintf(randomFormat(), name)
// 		a = append(a, msg)

// 	}
// return a,nil
// }
func Greet(names []string) ([]string, error) {
    var greetings []string
    for _, name := range names {
        if name == "" {
            return nil, errors.New("empty name")
        }
        msg := fmt.Sprintf(randomFormat(), name)
        greetings = append(greetings, msg)
    }
    return greetings, nil
}

	
// }
func randomFormat() string {
    // A slice of message formats.
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }
	return formats[rand.Intn(len(formats))]
}
