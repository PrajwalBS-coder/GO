package main


import(
	f "fmt"
	r "reflect"
)

func main(){
	s:=[]int{}
	f.Println(s,r.TypeOf(s))
	f.Println(len(s),cap(s))//here the length will be 0 and capacity will be 0


	s2:=[]int{1,2,3,4}
	f.Println(s2,r.TypeOf(s2))
	f.Println(len(s2),cap(s2))//here the length will be 4 and capacity will be 4

}