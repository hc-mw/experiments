package main

import (
	"encoding/json"
	"fmt"
	"unsafe"
)

type S struct {
	x int    `json:"x"`
	y string `json:"y"`
}

func main() {

	s := S{x: 42}
	ptr := &s // ptr points to s

	y := unsafe.Pointer(ptr)

	z := (*S)(y)
	p := *z
	fmt.Println(p)
	// print json for s
	q, err := json.MarshalIndent(p, "", "  ")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(q))

	// cmd := &cli.Command{
	// 	Name:  "greet",
	// 	Usage: "say a greeting",
	// 	Action: func(c *cli.Context) error {
	// 		if hostName, err := os.Hostname(); err == nil {
	// 			fmt.Printf("Hello %s\n", hostName)
	// 		} else {
	// 			log.Fatalln(err)
	// 		}
	// 		// fmt.Println("Greetings")
	// 		return nil
	// 	},
	// }

	// app := &cli.App{
	// 	Name:     "greet",
	// 	Usage:    "A demo cli app",
	// 	Commands: []*cli.Command{cmd},
	// }

	// if err := app.Run(os.Args); err != nil {
	// 	log.Fatalln(err.Error())
	// }
}
