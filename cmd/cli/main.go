package main

import (
	"fmt"
	"html/template"
	"os"
)

const (
	EXIT_SUCCESS int = 0
	EXIT_FAILED  int = -1
)

func run() int {
	fmt.Println("start")

	tmpl, err := template.ParseFiles("tmpl/README.gomd")
	if err != nil {
		fmt.Println("failed to load template file")
		return EXIT_FAILED
	}

	fd, err := os.Create("README.md")
	if err != nil {
		fmt.Println("failed to create README.md")
		return EXIT_FAILED
	}
	defer fd.Close()

	err = tmpl.Execute(fd, nil)
	if err != nil {
		fmt.Println("failed to write to README.md")
		return EXIT_FAILED
	}

	fmt.Println("end")
	return EXIT_SUCCESS
}

func main() {
	os.Exit(run())
}
