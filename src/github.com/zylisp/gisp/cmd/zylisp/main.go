package main

import (
	"github.com/zylisp/gisp/generator"
	"github.com/zylisp/gisp/repl"
	"os"
)

func main() {
	// XXX use cli opts/args here ...
	if len(os.Args) > 1 {
		generator.PrintGenerated(os.Args[1])
		return
	}
	repl.AstMain()
}
