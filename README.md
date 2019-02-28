# gisp

Simple (non standard) compiler of Lisp/Scheme to Go.

## Status

Project was largely abandoned in 2014, but revived in 2017 with updates from
various forks as well as some additional cleanup work.

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

## Usage

Start the REPL:

```
$ ./bin/gisp
```

From here you can type in forms and you'll get the Go AST back:

```lisp
>> (+ 1 1)
[(+ 1 1)]
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
    20  .  .  .  .  Name: "1"
    21  .  .  .  }
    22  .  .  }
    23  .  .  Ellipsis: -
    24  .  .  Rparen: -
    25  .  }
    26  }
```

To exit the REPL, just hit `<CONTROL>-C`.

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
