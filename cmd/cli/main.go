package main

import (
	"flag"
	"fmt"
	"html/template"
	"os"
)

const (
	EXIT_SUCCESS int = 0
	EXIT_FAILED  int = -1
)

func run(tmplPath string, output string) int {
	fmt.Println("start")

	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		fmt.Println("failed to load template file")
		return EXIT_FAILED
	}

	fd, err := os.Create(output)
	if err != nil {
		fmt.Printf("failed to create %s\n", output)
		return EXIT_FAILED
	}
	defer fd.Close()

	err = tmpl.Execute(fd, nil)
	if err != nil {
		fmt.Printf("failed to write to %s\n", output)
		return EXIT_FAILED
	}

	fmt.Println("end")
	return EXIT_SUCCESS
}

func main() {
	tmplPath := flag.String("tmpl_path", "tmpl/README.gomd", "path of README.md template file")
	output := flag.String("output", "README.md", "output file path")
	flag.Parse()

	os.Exit(run(*tmplPath, *output))
}
