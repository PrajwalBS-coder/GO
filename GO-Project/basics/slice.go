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


	//Creting slice through array
	s3:= []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //here data type will be []int here created variable using var keyword
	f.Println(s3, r.TypeOf(s3))
	p := s3[0:5] //here data type will be []int here created variable using var keyword
	f.Println(p, r.TypeOf(p))


	//Creting slice through array
	s4:=make([]float64,5,10) //here data type will be []int here created variable using var keyword
	f.Println(s4, r.TypeOf(s4))
	p1 := s4[0:5] //here data type will be []int here created variable using var keyword
	f.Println(p1, r.TypeOf(p1))



	//opertions on slice
	sl:=[]int{1,2,3,4,5,6,3}
	sl[0]=100//Changing values of perticular index
	f.Println(sl, r.TypeOf(sl))

	

}