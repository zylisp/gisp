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


## Installation

```bash
$ go get github.com/zylisp/zylisp/cmd/zylisp
```


## Development

```bash
$ git clone git@github.com:zylisp/zylisp.git
$ cd zylisp
$ . .env # optional, depending upon your local Go setup
$ make
```

That last step creates the `zylisp` binary and runs all the tests.

Note that the ZYLISP instructions and docs assume that the `git clone` has 
make made in the directory `github.com/zylisp` which is on the `GOPATH`.


## Usage

For usage as a REPL as well as a CLI, see the command documentation here:
 * https://zylisp.github.io/zylisp/doc/cmd/zylisp/

General package reference documentation is available here:
 * https://zylisp.github.io/zylisp/doc/


## Example Code

This is from the examples (all of which successfully compile from Lisp to both
Go source as well as bytecode):

```clj
(ns main
  "fmt"
  "github.com/zylisp/zylisp/core")

(def factorial (fn [n]
  (if (< n 2)
    1
    (* n (factorial (+ n -1))))))

(def main (fn []
  (fmt/printf "10! = %d\n"
              (int
                (assert
                  float64 (factorial 10))))))
```

See [examples](examples) for some more examples (they are Project Euler
solutions).


## Supported Lisp Functions

* `+`
* `-`
* `*`
* `mod`
* `let`
* `if`
* `ns`
* `def`
* `fn`
* All pre-existing Go functions


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
