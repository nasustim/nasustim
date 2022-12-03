package main

import (
	"fmt"
	"os"
)

const (
	EXIT_SUCCESS int = 0
	EXIT_FAILED  int = -1
)

func run() int {
	fmt.Println("start")

	fmt.Println("end")
	return EXIT_SUCCESS
}

func main() {
	os.Exit(run())
}
