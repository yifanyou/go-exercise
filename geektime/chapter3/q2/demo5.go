package main

import (
	"flag"
	"github.com/yifanyou/go-exercise/geektime/chapter3/q2/lib"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	lib.Hello(name)
}