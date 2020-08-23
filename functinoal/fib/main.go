package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"go_reset/functinoal/fib/fibs"
)

type intGen func() int

func (receiver intGen) Read(p []byte) (n int, err error) {
	next := receiver()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	var f intGen = fibs.Fibonacci()
	printFileContents(f)
}
