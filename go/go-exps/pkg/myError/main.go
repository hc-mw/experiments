package myerror

import (
	"errors"
	"fmt"
	"os"
)

func init() {
	fmt.Println("init 2")
}
func init() {
	fmt.Println("init")
}

// is and as in error
func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in fileChecker: %w", err)
	}
	f.Close()
	return nil
}

func ErrorDemo() {
	// err := fileChecker("not_here.txt")
	// if err != nil {
	// 	if errors.Is(err, os.ErrNotExist) {
	// 		fmt.Println("That file doesn't exist")
	// 	}
	// }

	a := 1
	b := 0
	x, err := divide(float64(a), float64(b))
	if err != nil {
		fmt.Printf(" %+v\n", err)
	} else {
		fmt.Println(x)
	}
}

func divide(a, b float64) (_ float64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("in divide: %+v", r)
		}
	}()

	if b == 0 {
		panic(errors.New("divide by zero"))
	}

	res := a / b

	return res, nil
}
