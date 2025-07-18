package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
)

func writeToWriter(w io.Writer, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		if pw, ok := w.(*io.PipeWriter); ok {
			pw.Close()
		}
	}()
	fmt.Fprintf(w, "Hello, writer\n")
}

func main() {
	// for {
	// 	rdr := io.LimitReader(os.Stdin, 4096)
	// 	_, err := io.Copy(os.Stdout, rdr)
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "Error copying: %s\n", err)
	// 		os.Exit(1)
	// 	}
	// 	// buf := make([]byte, 4096)
	// 	// n, err := rdr.Read(buf)
	// 	// if err != nil {
	// 	// 	fmt.Fprintf(os.Stderr, "Error reading: %s\n", err)
	// 	// 	os.Exit(1)
	// 	// }
	// 	// fmt.Fprintf(os.Stderr, "Read %d bytes, data: %s\n", n, string(buf[:n]))
	// }

	pr, pw := io.Pipe()

	var wg sync.WaitGroup

	wg.Add(1)

	go writeToWriter(pw, &wg)

	var b bytes.Buffer

	mw := io.MultiWriter(os.Stdout, &b)

	io.Copy(mw, pr)

	fmt.Printf("Buffer: %s\n", b.String())

	wg.Wait()
}
