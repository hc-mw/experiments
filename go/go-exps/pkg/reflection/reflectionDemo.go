package reflection

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
)

func printReflection(obj any) {
	rType := reflect.TypeOf(obj)
	rVal := reflect.ValueOf(obj)

	fmt.Printf("-------PRINTING REFLECTION FOR \"%v\"------\n", rVal)
	fmt.Println("reflect type=>", rType, ", reflect value=>", rVal.String())
	fmt.Println("reflect.Value.Type() method: ", rVal.Type())
	fmt.Println("reflect.Value.Kind() method: ", rVal.Kind())
	fmt.Println("reflect.Type.Kind() method: ", rType.Kind())
	if reflect.Uint64 == rType.Kind() {
		fmt.Println("reflect.Type.Uint() method: ", rVal.Uint())
	}
	if reflect.Float64 == rType.Kind() {
		fmt.Println("reflect.Type.Float() method: ", rVal.Float())
	}
	fmt.Println("---------------------------------------")
}

func reflectionSetDemo() {
	var pi float64 = 3.14
	v := reflect.ValueOf(pi)
	fmt.Println(v.CanSet()) // false
	v = reflect.ValueOf(&pi)
	fmt.Println(v.CanSet()) // false
	p := v.Elem()
	fmt.Println(p.CanSet()) // true
	if p.CanSet() {
		p.SetFloat(3.1)
	}
	fmt.Println(pi) // 3.1
}

func ReflectionDemo() {
	var r io.Reader
	r = bytes.NewBuffer([]byte("Hey there"))
	if _, ok := r.(io.Reader); ok {
		fmt.Println("r implements io.Reader")
	}
	printReflection(r)

	var pi float64 = 3.14
	printReflection(pi)

	var x any
	x = 12
	printReflection(x)

	iterateOnStruct()
}

func iterateOnStruct() {
	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}
	t := Person{"john", "doe", 23}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	if s.Kind() == reflect.Struct {
		for i := range s.NumField() {
			f := s.Field(i)
			fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
		}
	}

	s.Field(0).SetString("johny")
	fmt.Println("t is now ", t)
}
