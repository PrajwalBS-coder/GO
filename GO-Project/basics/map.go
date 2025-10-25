// Go Maps
// Maps are used to store data values in key:value pairs.

// Each element in a map is a key:value pair.

// A map is an unordered and changeable collection that does not allow duplicates.

// The length of a map is the number of its elements. You can find it using the len() function.

// The default value of a map is nil.

// Maps hold references to an underlying hash table.

// Go has multiple ways for creating maps.
// Syntax
// var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
// b := map[KeyType]ValueType{key1:value1, key2:value2,...}

package main

import (
	f "fmt"
	r "reflect"
)
func main() {
	var m map[string]int //here data type will be map[string]int here created variable using var keyword
	f.Println(m, r.TypeOf(m))
	m = make(map[string]int) //here data type will be map[string]int here created variable using make keyword
	f.Println(m, r.TypeOf(m))
	m["key"] = 10 //here data type will be map[string]int here created variable using var keyword
	m["value"] = 20
	f.Println(m, r.TypeOf(m))
	f.Println(m["key"], r.TypeOf(m["key"])) //here data type will be int here created variable using var keyword

	a:=map[string]int{"key":10,"value":20}
	f.Println(a, r.TypeOf(a))

//Syntax for map with make
// m := make(map[string]int) within [] is key and int is value


	b:=make(map[string]int)//Create empty map with use of make
	b["key"] = 10
	b["value"] = 20
	f.Println(b, r.TypeOf(b))


	// Allowed Key Types

// 	The map key can be of any data type for which the equality operator (==) is defined. These include:

// Booleans
// Numbers
// Strings
// Arrays
// Pointers
// Structs
// Interfaces (as long as the dynamic type supports equality)
// Invalid key types are:

// Slices
// Maps
// Functions
//These types are invalid because the equality operator (==) is not defined for them.

 var p = make(map[string]string)
  p["brand"] = "Ford"
  p["brand"] = "Ford"
  p["model"] = "Mustang"
  p["year"] = "1964"

  f.Printf(p["brand"])
  //map_name[key] = value
  p["year"] = "1970" // Updating an element
  p["color"] = "red" // Adding an element

  f.Println(p)

// Remove Element from Map
// Removing elements is done using the delete() function.

// Syntax
// delete(map_name, key)

delete(p, "model")
f.Println(p)


// Check For Specific Elements in a Map
// Syntax
// val, ok :=map_name[key]

val, ok := p["VBN"]
f.Println(val, ok)//val actaull value of year if the key is there and ok will be true else it will be false

var l =make(map[string]int)
f.Println(l, r.TypeOf(l))

f.Println(l["key"])


// Maps are references to hash tables.

// If two map variables refer to the same hash table, changing the content of one variable affect the content of the other.

var k=p
k["year"] = "1190"
f.Println(k, r.TypeOf(k))
f.Println(p, r.TypeOf(p))

// Iterate Over Maps
// You can use range to iterate over maps.

for key,value := range p{
	f.Println(key, value)
}

// Iterate Over Maps in a Specific Order
// Maps are unordered data structures. If you need to iterate over a map in a specific order, you must have a separate data structure that specifies that order.


 a= map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

  var x []string             // defining the order
  x = append(x, "one", "two", "three", "four")

  for k, v := range a {        // loop with no order
    f.Printf("%v : %v, ", k, v)
  }

  f.Println()

  for _, element := range x {  // loop with the defined order
    f.Printf("%v : %v, ", element, a[element])
  }

}