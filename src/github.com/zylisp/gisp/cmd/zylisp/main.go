package main

import (
  "flag"
  "fmt"
	"github.com/zylisp/gisp/generator"
	"github.com/zylisp/gisp/repl"
	"os"
)

func main() {
	astPtr := flag.Bool("ast", false, "Enable AST mode")
	cliPtr := flag.Bool("cli", false, "Run as a CLI tool")
	// dirPtr := flag.String("dir", "/tmp", "Default directory for writing operations")
	goPtr := flag.Bool("go", false, "Enable Go code-generation mode")
	lispPtr := flag.Bool("lisp", false, "Enable LISP mode")
	// outPtr := flag.String("o", "output", "Default filename for writing operations")

	flag.Parse()
	files := flag.Args()
	if *cliPtr == true && len(files) < 1 {
		fmt.Println("You need to provide at least one file upon which to operate")
		os.Exit(1)
	}
	if *lispPtr == true {
		if *cliPtr == true {
			fmt.Println("The Lisp CLI is currently not supported")
		} else if *cliPtr == false {
			fmt.Println("Lisp mode is currently not supported")
		}
	} else if *astPtr == true {
		if *cliPtr == true {
			fmt.Println("AST CLI is not currently supported")
		} else if *cliPtr == false {
			repl.AstMain()
		}
	} else if *goPtr == true {
		if *cliPtr == true {
			for _, file := range files {
				// XXX check to see if printing or saving to file
				generator.PrintGenerated(file)
			}
		} else if *cliPtr == false {
			fmt.Println("Go mode is currently not supported")
		}
	} else {
		fmt.Println("You need to supply a mode")
	}
}
