package main

import f "fmt"
import r "reflect"

func main() {
	var a int = 1000
	f.Println(a, r.TypeOf(a))
	switch a {
	case 10:
		f.Println("A is 10")
	case 20:
		f.Println("A is 20")
	default:
		f.Println("A is not 10 or 20")
	}
}