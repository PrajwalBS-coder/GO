package main

import (
	"fmt"
	"log"

	"amin.com/temp"
	"gonum.org/v1/gonum/stat"
)

func main(){
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	l:=[] float64{1, 2, 3, 4, 5}
	mean := stat.Mean(l, nil)
	fmt.Println("Mean:", mean)
	median := stat.Quantile(0.5, stat.Empirical, l, nil)
	fmt.Println("Median:", median)
	name:=[]string{"AMin","Jhon"}
	msg,err:=temp.Greet(name)
	if err!=nil{
		fmt.Println(err)
		log.Fatal(err)
	}else{
		fmt.Println(msg)
	}
}