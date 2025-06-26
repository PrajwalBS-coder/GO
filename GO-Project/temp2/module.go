package main

import (
	"fmt"
"gonum.org/v1/gonum/stat"
"amin.com/temp"
)

func main(){
	l:=[] float64{1, 2, 3, 4, 5}
	mean := stat.Mean(l, nil)
	fmt.Println("Mean:", mean)
	median := stat.Quantile(0.5, stat.Empirical, l, nil)
	fmt.Println("Median:", median)
	temp.Greet("John")
}