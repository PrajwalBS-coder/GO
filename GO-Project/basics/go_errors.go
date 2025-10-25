package main

import (
    "errors"
    "fmt"
)

func f(arg int) (int, error) {
    if arg == 42 {

        return -1, errors.New("can't work with 42")
    }

    return arg + 3, nil
}

var ErrOutOfTea = errors.New("no more tea available")
var ErrPower = errors.New("can't boil water")

func makeTea(arg int) error {
    if arg == 2 {
        return ErrOutOfTea
    } else if arg == 4 {

        return fmt.Errorf("making tea: %w", ErrPower)
    }
    return nil
}

func main() {
    for _, i := range []int{7, 42} {

        if r, e := f(i); e != nil {
            fmt.Println("f failed:", e)
        } else {
            fmt.Println("f worked:", r)
        }
    }

    for i := range 5 {
        if err := makeTea(i); err != nil {

            if errors.Is(err, ErrOutOfTea) {
                fmt.Println("We should buy new tea!")
            } else if errors.Is(err, ErrPower) {
                fmt.Println("Now it is dark.")
            } else {
                fmt.Printf("unknown error: %s\n", err)
            }
            continue
        }

        fmt.Println("Tea is ready!")
    }
}

// In Go it’s idiomatic to communicate errors via an explicit, separate return value. 
// This contrasts with the exceptions used in languages like Java, Python and Ruby and the overloaded single result / error value sometimes used in C. 
// Go’s approach makes it easy to see which functions return errors and to handle them using the same language constructs employed for other, non-error tasks.
// See the documentation of the errors package and this blog post for additional details.

