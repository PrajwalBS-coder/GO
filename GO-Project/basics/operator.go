package main

import (
	f "fmt"
	r "reflect"
)

func main() {
	//arithmatic operations

	var a float32 = 10 + 20
	f.Println(a, r.TypeOf(a))
	b := 10 - 20
	f.Println(b, r.TypeOf(b))
	c := 10 * 20
	f.Println(c, r.TypeOf(c))
	d := 10 / 20
	f.Println(d, r.TypeOf(d))
	e := 10 % 20
	f.Println(e, r.TypeOf(e))
	g := 90
	g++
	f.Println(g, r.TypeOf(g))
	g--
	f.Println(g, r.TypeOf(g))

	//assignment operations

	var x int = 10
	f.Println(x, r.TypeOf(x))
	x += 20
	f.Println(x, r.TypeOf(x))
	x -= 20
	f.Println(x, r.TypeOf(x))
	x *= 20
	f.Println(x, r.TypeOf(x))
	x /= 20
	f.Println(x, r.TypeOf(x))
	x %= 20
	f.Println(x, r.TypeOf(x))
	x &= 3
	f.Println(x, r.TypeOf(x))
	x |= 3
	f.Println(x, r.TypeOf(x))
	x ^= 2
	f.Println(x, r.TypeOf(x))
	x <<= 3
	f.Println(x, r.TypeOf(x))
	x >>= 3
	f.Println(x, r.TypeOf(x))
	

	//Comparison operator

var	z=10
var y=20
f.Println(z==y)
f.Println(z!=y)
f.Println(z>y)
f.Println(z<y)
f.Println(z>=y)
f.Println(z<=y)


//Logicl opertor
op:=10
f.Println(op>1 && op==10)
f.Println(op>10 || op==10)
f.Println(op !=10)

//Bitwise Operators

f.Println(10&9)//1010 +1001==1000=>8
f.Println(10|9)//1010 +1001==1000=>11
f.Println(10^11)//1010+1011==0001=>1
f.Println(10<<1)//1010-->10100=>20
f.Println(10>>1)//1010-->0101=>5

}
