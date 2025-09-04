package main

import f "fmt"
import r "reflect"

func main(){
var  a float32=10+20
f.Println(a,r.TypeOf(a))
b:=10-20
f.Println(b,r.TypeOf(b))
c:=10*20
f.Println(c,r.TypeOf(c))
d:=10/20
f.Println(d,r.TypeOf(d))
e:=10%20
f.Println(e,r.TypeOf(e))
g:=90
g++
f.Println(g,r.TypeOf(g))
g--
f.Println(g,r.TypeOf(g))

}