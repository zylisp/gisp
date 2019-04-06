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
make made in the directory `src/github.com/zylisp`, where the parent directory
of `src` is on the `GOPATH`.


## Usage

For usage as a REPL as well as a CLI, see the command documentation here:
 * https://zylisp.github.io/zylisp/doc/cmd/zylisp/

 In short, once compiled, you may pass a flag for one of the supported REPL
 modes (e.g., `-ast`), or use `zylisp` as a CLI tool (i.e., compiler), with the 
 `-cli` flag.

General package reference documentation is available here:
 * https://zylisp.github.io/zylisp/doc/


## Docker Support

For those who have `docker` installed and do not wish to install Go, you may 
try out the various REPLs via `docker` commands, e.g.:

```bash
docker run -it zylisp/zylisp:latest -ast
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

ZYLISP version: 1.0.0-alpha4
Build: release/1.0.x@120e6d5, 2019-04-06T21:59:42Z
REPL Mode: AST
Go version: go1.12

        Docs: https://zylisp.github.io/zylisp/
     Project: https://github.com/zylisp/zylisp
Instructions: Simply type any form to view the generated Go AST.
        Exit: ^D or ^C

AST>
```

Futhermore, since `zylisp` is the entrypoint for the Docker image, the run 
command may receive all the options that the `zylisp` binary receives, 
including the help flag:

```bash
$ docker run -it zylisp/zylisp:latest -h
```
```
Usage of zylisp:
  -ast
    	Enable AST mode
  -bytecode
    	Enable byte-code compilation from generated Go
  -cli
    	Run as a CLI tool
  -dir string
    	Default directory for writing operations
  -go
    	Enable Go code-generation mode
  -lisp
    	Enable LISP mode
  -loglevel string
    	Set the logging level (default "warning")
  -o string
    	Default filename for writing operations
  -version
    	Display version/build info and exit
```

Note that the ZYLISP docker images are very small, usually weighing in about
4 MB in size.


## Credits

* @jcla1 for the initial implementation
* @masukomi for adding a number of tests and checks
* @m90 for README fixes
* The ZYLISP project for new development


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
