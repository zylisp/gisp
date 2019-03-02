/*
The ZYLISP command line and multi-REPL tool

Overview

TBD

CLI Mode

TBD

REPL Mode

TBD

*/
package main

import (
  "flag"
  "fmt"
	"github.com/zylisp/gisp/generator"
	"github.com/zylisp/gisp/repl"
	"os"
)

func dispatchLisp(isCli bool) {
	if isCli {
		// LISP CLI
		fmt.Println("The Lisp CLI is currently not supported")
	} else {
		// LISP REPL
		fmt.Println("Lisp mode is currently not supported")
		// repl.LispMain()
	}
}

func dispatchAst(isCli bool, files []string) {
	if isCli {
		// AST CLI
		for _, file := range files {
			// XXX check to see if printing or saving to file; currently
			generator.PrintGeneratedAst(file)
		}
	} else {
		// AST REPL
		repl.AstMain()
	}
}

func dispatchGoGen(isCli bool, files []string) {
	if isCli {
		// Go-generator CLI
		for _, file := range files {
			// XXX check to see if printing or saving to file
			generator.PrintGeneratedGo(file)
		}
	} else {
		// Go-generator REPL
		fmt.Println("Go-gen mode is currently not supported")
		// repl.GoGenMain()
	}
}

func main() {
	astPtr := flag.Bool("ast", false, "Enable AST mode")
	cliPtr := flag.Bool("cli", false, "Run as a CLI tool")
	// dirPtr := flag.String("dir", "/tmp", "Default directory for writing operations")
	goPtr := flag.Bool("go", false, "Enable Go code-generation mode")
	lispPtr := flag.Bool("lisp", false, "Enable LISP mode")
	// outPtr := flag.String("o", "output", "Default filename for writing operations")

	flag.Parse()
	files := flag.Args()
	if *cliPtr && len(files) < 1 {
		fmt.Println("You need to provide at least one file upon which to operate")
		os.Exit(1)
	}
	if *lispPtr {
		dispatchLisp(*cliPtr)
	} else if *astPtr {
		dispatchAst(*cliPtr, files)
	} else if *goPtr {
		dispatchGoGen(*cliPtr, files)
	} else {
		fmt.Println("You need to supply a mode")
	}
}
