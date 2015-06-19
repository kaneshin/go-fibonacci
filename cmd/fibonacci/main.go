package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/kaneshin/go-fibonacci/fibonacci"
)

func usage(io *os.File) {
	fmt.Fprintf(io, `fibonacci

  Usage:
    fibonacci -h/--help
    fibonacci [number]
`)
}

func main() {
	flag.Usage = func() {
		usage(os.Stdout)
	}
	flag.Parse()

	if flag.NArg() == 0 {
		os.Stdin
		os.Exit(0)
	}

	arg := flag.Arg(0)
	num, err := strconv.Atoi(arg)
	if err != nil {
		usage(os.Stderr)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "%d\n", fibonacci.Fibonacci(num))
}
