package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {
	verbose := flag.Bool("v", false, "verbose output")
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("missing path to file")
		fmt.Println("")
		fmt.Println("usage: jammel [-v] <path>")
		os.Exit(1)
	}

	path := args[0]

	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("failed to open file '%s': %v\n", path, err)
		os.Exit(1)
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Printf("failed to read file '%s': %v\n", path, err)
		os.Exit(1)
	}

	var res map[string]interface{}
	err = yaml.Unmarshal(b, &res)
	if err != nil {
		fmt.Printf("failed to parse file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("successfully parsed file")

	if *verbose {
		fmt.Printf("%v\n", res)
	}
}
