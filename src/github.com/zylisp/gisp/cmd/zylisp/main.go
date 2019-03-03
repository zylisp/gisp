/*
The ZYLISP command line and multi-REPL tool.

Overview

The ZYLISP project's zylisp command is a tool for both performing command line
actions (i.e., batch jobs) as well as interactive programming (REPLs). In both
cases, there are three modes:

 * AST
 * Go code generation
 * Lisp (the actual, classic REPL)

Each of these is covered in more detail below.

Setup

For setting the project GOPATH and building the zylisp command, see the project
Github page at https://github.com/zylisp/zylisp#development.

If you would prefer to use ZYLISP without cloning the repo and setting up a
development environment, you may install it with the following:

	$ go get github.com/zylisp/gisp/cmd/zylisp

Since the zylisp command makes use of Go flags, it has a generated help output.
You may view this with the usual -h option:

	$ zylisp -h

REPL

The ZYLISP REPL is designed for interactive programming.

In AST mode, the REPL will take any expression given and attempt to create
a Go abstract syntax tree from that:

	$ zylisp -ast

	AST> (+ 1 2)
	Parsed:
	[(+ 1 2)]
	AST:
	     0  []ast.Expr (len = 1) {
	     1  .  0: *ast.CallExpr {
	     2  .  .  Fun: *ast.SelectorExpr {
	     3  .  .  .  X: *ast.Ident {
	     4  .  .  .  .  NamePos: -
	     5  .  .  .  .  Name: "core"
	     6  .  .  .  }
	     7  .  .  .  Sel: *ast.Ident {
	     8  .  .  .  .  NamePos: -
	     9  .  .  .  .  Name: "ADD"
	    10  .  .  .  }
	    11  .  .  }
	    12  .  .  Lparen: -
	    13  .  .  Args: []ast.Expr (len = 2) {
	    14  .  .  .  0: *ast.Ident {
	    15  .  .  .  .  NamePos: -
	    16  .  .  .  .  Name: "1"
	    17  .  .  .  }
	    18  .  .  .  1: *ast.Ident {
	    19  .  .  .  .  NamePos: -
	    20  .  .  .  .  Name: "2"
	    21  .  .  .  }
	    22  .  .  }
	    23  .  .  Ellipsis: -
	    24  .  .  Rparen: -
	    25  .  }
	    26  }


CLI

The command line interface is for performing batched or scripted operations
and is enabled with the -cli option. However, in CLI mode, you must pass one
or more file names (file globbing allowed).

In AST mode, the CLI will parse the given file and produce as output the
corresponding Go AST for the Lisp code in the file provided:

	$ zylisp -ast -cli examples/factorial.gsp

The output is a bit long (302 lines), but here are the first 20 lines:

     0  *ast.File {
     1  .  Package: -
     2  .  Name: *ast.Ident {
     3  .  .  NamePos: -
     4  .  .  Name: "main"
     5  .  }
     6  .  Decls: []ast.Decl (len = 3) {
     7  .  .  0: *ast.GenDecl {
     8  .  .  .  TokPos: -
     9  .  .  .  Tok: import
    10  .  .  .  Lparen: -
    11  .  .  .  Specs: []ast.Spec (len = 2) {
    12  .  .  .  .  0: *ast.ImportSpec {
    13  .  .  .  .  .  Path: *ast.BasicLit {
    14  .  .  .  .  .  .  ValuePos: -
    15  .  .  .  .  .  .  Kind: STRING
    16  .  .  .  .  .  .  Value: "\"fmt\""
    17  .  .  .  .  .  }
    18  .  .  .  .  .  EndPos: -
    19  .  .  .  .  }
    20  .  .  .  .  1: *ast.ImportSpec {

In Go-generation mode, the output is instead compilable Go code:

	$ zylisp -cli -go examples/factorial.gsp

	package main

	import (
		"fmt"
		"github.com/zylisp/gisp/core"
	)

	func main() {
		fmt.Printf("10! = %d\n", int(factorial(10).(float64)))
	}
	func factorial(n core.Any) core.Any {
		return func() core.Any {
			if core.LT(n, 2) {
				return 1
			} else {
				return core.MUL(n, factorial(core.ADD(n, -1)))
			}
		}()
	}


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

func dispatchAST(isCli bool, files []string) {
	if isCli {
		// AST CLI
		for _, file := range files {
			// XXX check to see if printing or saving to file; currently
			generator.PrintASTFromFile(file)
		}
	} else {
		// AST REPL
		repl.ASTMain()
	}
}

func dispatchGoGen(isCli bool, files []string) {
	if isCli {
		// Go-generator CLI
		for _, file := range files {
			// XXX check to see if printing or saving to file
			generator.PrintGoFromFile(file)
		}
	} else {
		// GOGEN REPL
		repl.GoGenMain()
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
		dispatchAST(*cliPtr, files)
	} else if *goPtr {
		dispatchGoGen(*cliPtr, files)
	} else {
		fmt.Println("You need to supply a mode")
	}
}
