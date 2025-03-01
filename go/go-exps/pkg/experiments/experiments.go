package experiments

import "fmt"

func StringDemo() {
	str := "Hello there"

	fmt.Printf("type of str[1]: %T\n", str[1])
	fmt.Printf("type of str[1]: %v\n", str[1])

	str1 := str[1:4]
	fmt.Printf("type of str1: %T\n", str1)
	fmt.Printf(" str1: %v\n", str1)

	var a rune = 'x'
	var b = string(a)
	fmt.Printf("type of b: %v\n", b)
}

func StructDemo() {
	// anonymous struct
	var person struct {
		name string
		age  int
		pet  string
	}

	person.name = "bob"
	person.age = 50
	person.pet = "dog"

	fmt.Println(person)

	pet := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "dog",
	}

	fmt.Println(pet)

	type firstPerson struct {
		name string
		age  int
	}
	f := firstPerson{
		name: "Bob",
		age:  50,
	}

	var g struct {
		name string
		age  int
	}

	// compiles -- can use = and == between identical named and anonymous structs
	g = f
	fmt.Println(f == g)
}

func IteratingDemo() {
	samples := []string{"hello", "apple_π!"}
outer:
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'l' {
				continue outer
			}
		}
		fmt.Println()
	}
}
