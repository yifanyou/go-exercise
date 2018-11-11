package main

import (
	"fmt"
	"flag"
	"os"
)

var name string

func init() {
	//flag.CommandLine = flag.NewFlagSet("", flag.PanicOnError)
	//flag.CommandLine.Usage = func() {
	//	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	//	flag.PrintDefaults()
	//}
	//
	//flag.StringVar(&name, "name", "everyone", "The greeting object.")

	cmdLine.StringVar(&name, "name", "everyone", "The greeting object.")
}

var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)

func main() {
	//flag.Usage = func() {
	//	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	//	flag.PrintDefaults()
	//}

	//flag.Parse()

	cmdLine.Parse(os.Args[1:])

	fmt.Printf("Hello, %s\n", name)
}