package main

import (
	f "fmt" 
	r "reflect"
)
func main(){
	var array=[]int{1,2}
	f.Println(array,r.TypeOf(array))
	a:=[]int{2,3}
	f.Println(a,r.TypeOf(a))
}