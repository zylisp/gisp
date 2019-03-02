# ZYLISP

*Simple (non standard) compiler of Lisp/Scheme to Go*

[![Build Status][travis-badge]][travis]
[![Tag][tag-badge]][tag]
[![Go version][go-v]](.travis.yml)


## Status

This project was largely abandoned in 2014, but revived in 2017 with updates
from various forks as well as some additional cleanup work. It was brought into
the ZYLISP Github org for exploratory purposes, and received more loving tweaks
and cleanups. More to come ...


## Includes

- Lexer based on Rob Pike's
  [Lexical Scanning in Go](https://talks.golang.org/2011/lex.slide)
- Simple recursive parser, supporting ints, floats, strings, bools
- TCO via loop/recur
- AST generating REPL included


## Development

```bash
$ git clone git@github.com:zylisp/gisp.git
$ cd gisp
$ export GOPATH=$GOPATH:`pwd`
$ export PATH=$PATH:`pwd`/bin
$ make
```

That last step creates the `zylisp` binary and runs all the tests.

Note that the ZYLISP instructions and docs assume the `./bin` dir has been
added to the `PATH` as above.


## Usage

### AST REPL

Start the REPL by executing the binary with the appropriate flag:
```
$ ./bin/zylisp -ast
```
```
Okay, 3, 2, 1 - Let's jam!

Welcome to

/^^^^^^^^/^^ /^^      /^^ /^^       /^^ /^^^^^^^^ /^^^^^^^^^
       /^^    /^^    /^^  /^^       /^^ /^^       /^^    /^^
      /^^      /^^ /^^    /^^       /^^ /^^       /^^    /^^
    /^^          /^^      /^^       /^^ /^^^^^^^^ /^^^^^^^^^
   /^^           /^^      /^^       /^^       /^^ /^^
 /^^             /^^      /^^       /^^ /^^   /^^ /^^
/^^^^^^^^^^^     /^^      /^^^^^^^^ /^^ /^^^^^^^^ /^^

ZYLISP version: 0.7.0-alpha1/47d1949 [AST mode]
Go version: go1.12

        Docs: https://zylisp.github.io/zylisp/
     Project: https://github.com/zylisp/zylisp
Instructions: Simply type any form to view the generated Go AST.
        Exit: <CONTROL><C>

AST>
```

From here you can type in forms and you'll get the Go AST back:

```lisp
AST> (do-something 1 2)
```
```
Parsed:
[(do-something 1 2)]
AST:
     0  []ast.Expr (len = 1) {
     1  .  0: *ast.CallExpr {
     2  .  .  Fun: *ast.Ident {
     3  .  .  .  NamePos: -
     4  .  .  .  Name: "doSomething"
     5  .  .  }
     6  .  .  Lparen: -
     7  .  .  Args: []ast.Expr (len = 2) {
     8  .  .  .  0: *ast.Ident {
     9  .  .  .  .  NamePos: -
    10  .  .  .  .  Name: "1"
    11  .  .  .  }
    12  .  .  .  1: *ast.Ident {
    13  .  .  .  .  NamePos: -
    14  .  .  .  .  Name: "2"
    15  .  .  .  }
    16  .  .  }
    17  .  .  Ellipsis: -
    18  .  .  Rparen: -
    19  .  }
    20  }
```

To exit the REPL, just hit `<CONTROL><C>`.


### Go REPL

TBD


### Lisp REPL

TBD


### CLI

You may also call `zylisp` as a command line tool by passing the `cli` flag.
Currently, only the following command-line modes are supported:

* `zylisp -go -cli`

To generate Go code from a Lisp file:

```
$ zylisp -go -cli examples/factorial.gsp
```
```go
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
```

You may also generate code for multiple files at once using file system globs:

```
$ zylisp -go -cli examples/*.gsp
```

At a future date writing these to files will also be supported.


### Compiler

To compile:

```
$ zyc examples/even_fib_terms.gsp
```

By default, this will create the executable binary file `even_fib_terms` in
the current working directory. You also have the option of specifying the
filename/path of the output:

```
$ zyc -o bin/examples/fib-even-terms examples/even_fib_terms.gsp
```

Or, if you prefer, you can compile all the `.gsp` files in a directory, in
which case the output option is interpreted as a directory:

```
$ zyc -o bin/examples examples/*.gsp
```

Then you can run them on your system as any compiled Go:

```
$ ./bin/examples/sum-fib-terms
Sum of all even fibonacci terms below 4000000: 4613732
```

To see the other compiler options available, run `zyc -h`.

Note that the compilation process involves parsing, generating an AST,
generating Go code, and finally, compiling that Go code.


## Supported Lisp Functions

```
+, -, *, mod, let, if, ns, def, fn, all pre-existing Go functions
```

See [examples](examples) for some Project Euler solutions


## Credits

* @jcla1 for the initial implementation
* @masukomi for adding a number of tests and checks
* @m90 for README fixes


## License

MIT


<!-- Named page links below: /-->

[logo]: media/images/logo-1-250x.png
[logo-large]: media/images/logo-1.png
[travis]: https://travis-ci.org/zylisp/zylisp
[travis-badge]: https://travis-ci.org/zylisp/zylisp.png?branch=master
[tag-badge]: https://img.shields.io/github/tag/zylisp/zylisp.svg
[tag]: https://github.com/zylisp/zylisp/tags
[go-v]: https://img.shields.io/badge/Go-1.12-blue.svg
