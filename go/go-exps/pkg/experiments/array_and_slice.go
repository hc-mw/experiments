package experiments

import (
	"fmt"
	"strings"
)

func ArrayDemo() {
	p := [4]int{1, 2, 3, 4}

	fmt.Println(cap(p), len(p))

	t := p[:]

	capAndLen("t", t)

	x := make([]int, 0, 100)

	for i := 0; i < 4; i++ {
		x = append(x, i+1)
	}

	y := x[2:5]

	z := x[2:5:5]

	capAndLen("x", x)
	capAndLen("y", y)
	capAndLen("z", z)
	fmt.Println("x: ", x)
	fmt.Println("y: ", y)
	fmt.Println("z: ", z)

	y = append(y, 30)
	y = append(y, 40)
	z = append(z, 50)

	fmt.Println(strings.Repeat("=", 20))

	capAndLen("x", x)
	capAndLen("y", y)
	capAndLen("z", z)
	fmt.Println("x: ", x)
	fmt.Println("y: ", y)
	fmt.Println("z: ", z)

	fmt.Println(strings.Repeat("=", 20))

	copy(x[:3], x[1:])
	fmt.Println("x: ", x)
}

func capAndLen(x string, arr []int) {
	fmt.Printf("%s Capacity: %d, Length: %d\n", x, cap(arr), len(arr))
}
