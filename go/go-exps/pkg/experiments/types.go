package experiments

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) a() {
	fmt.Println(p)
}

func (p *Person) b() {
	fmt.Println(p.Name)
}

func NewPerson(name string, age int) Person {
	return Person{
		Name: name,
		Age:  age,
	}
}

// function types
type Logger interface {
	Logger(x string)
}

type LoggerAdapter func(message string)

func (la LoggerAdapter) Log(message string) {
	la(message)
}

func LogOutput(message string) {
	fmt.Println(message)
}

func nullInstanceDemo() {
	p1 := NewPerson("bob", 50)
	// t := NewPerson("alex", 45)
	var p2 *Person

	if p2 == nil {
		fmt.Println("p2 is nil")
	}

	p1.a()
	p1.b()

	// p2.a()
	p2.b()
}

func TypeDemo() {
	la := LoggerAdapter(LogOutput)
	la("hello")
}
