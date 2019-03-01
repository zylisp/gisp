# gisp

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


## Usage

Start the REPL by executing the binary:

```
$ ./bin/zylisp
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

To generate Go:

```
$ ./bin/gisp examples/even_fib_terms.gsp > examples/even_fib_terms.go
```

To compile the generated Go:

```
$ go build -o ./bin/sum-fib-terms examples/even_fib_terms.go
```

Then you can run on your system:

```
$ ./bin/sum-fib-terms
Sum of all even fibonacci terms below 4000000: 4613732
```

## Functions

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
[travis]: https://travis-ci.org/zylisp/gisp
[travis-badge]: https://travis-ci.org/zylisp/gisp.png?branch=master
[tag-badge]: https://img.shields.io/github/tag/zylisp/gisp.svg
[tag]: https://github.com/zylisp/gisp/tags
[go-v]: https://img.shields.io/badge/Go-1.12-blue.svg
