package main

import (
	"github.com/yukpiz/go-module-example/ex1"
	"github.com/yukpiz/go-module-example/ex2"
	"github.com/yukpiz/go-module-example/nst/ex3"
	"github.com/yukpiz/go-module-example/other_branch/ex4"
)

func main() {
	ex1.Ex1Hello() // ex1 hello!
	ex2.Ex2Hello() // ex2 hello!
	ex3.Ex3Hello() // ex3 hello!
	ex4.Ex4Hello() // ex4 hello!
}
