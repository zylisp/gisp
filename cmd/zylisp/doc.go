/*
The ZYLISP command line and multi-REPL tool.

Overview

The ZYLISP project's zylisp command is a tool for both performing command line
actions (i.e., batch jobs) as well as interactive programming (REPLs). In both
cases, there are three modes:

   * AST
   * Go code generation
   * Byte-code compilation (no REPL; just CLI support)
   * Lisp (the actual, classic REPL)

Each of these is covered in more detail below.

Setup

For setting the project GOPATH and building the zylisp command, see the project
Github page at https://github.com/zylisp/zylisp#development.

If you would prefer to use ZYLISP without cloning the repo and setting up a
development environment, you may install it with the following:

	$ go get github.com/zylisp/zylisp/cmd/zylisp

Since the zylisp command makes use of Go flags, it has a generated help output.
You may view this with the usual -h option:

	$ zylisp -h

Logging

The zylisp executable supports passing a -loglevel option with one of the
following a legal associated value:

  * debug
  * info
  * notice
  * warning
  * error
  * failure


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

	$ zylisp -ast -cli examples/factorial.zsp

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

	$ zylisp -cli -go examples/factorial.zsp

	package main

	import (
		"fmt"
		"github.com/zylisp/zylisp/core"
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

In byte-code compilation mode, Go is generated under the hood, and then it is
compiled to byte code using 'go build':

  $ zylisp -cli -bytecode -dir bin/examples examples/*.zsp

This also demonstrates support for file globbing, allowing you to generate
output for multiple files at once.

For convenience, the 'zyc' bash wrapper is provided for
'zylisp -cli -bytecode':

  $ zyc -dir bin/examples examples/*.zsp

Note that since 'zyc' is not a compiled command, it will not be installed by
Go. A 'make install-zyc' target is provided, however, that will install `zyc`
into '~/go/bin'.

*/
package main
