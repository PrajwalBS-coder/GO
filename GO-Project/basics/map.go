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
import f "fmt"
import r "reflect"
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
}