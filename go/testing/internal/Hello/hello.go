package hello

import "fmt"

func Hello(name string) string {
	//if name == "" {
	//	name = "world"
	//}
	return fmt.Sprintf("Hello %s", name)
}

func main() {
	fmt.Println(Hello("world"))
}
