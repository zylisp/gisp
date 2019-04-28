# ZYLISP

[![Build Status][travis-badge]][travis]
[![Tag][tag-badge]][tag]
[![Go version][go-v]](.travis.yml)

[![][logo]][logo-large]

_A Simple Lisp that compiles to Go_

## Status

ZYLISP is capable of generating Go files and compiled byte code from files
written in the ZYLISP Lisp dialect. AST files may also be generated from the
command line. An AST-generating expression shell is currently available, but
no LISP REPL yet.

Current development efforts are focused on adding more core functions to the
Lisp dialect with a special interest in syntactic support for explicit types.

Milestone currently under development:
[0.9.0](https://github.com/zylisp/zylisp/milestone/4)

## Features

- Simple recursive parser, supporting ints, floats, strings, bools
- TCO via loop/recur
- AST-generating shell and CLI
- Go-generating CLI
- Byte-code-compiling CLI
- Published Docker images

## Supported Lisp Functions

- `+`
- `-`
- `*`
- `mod`
- `let`
- `if`
- `ns`
- `def`
- `fn`
- All pre-existing Go functions

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

- [https://zylisp.github.io/zylisp/doc/cmd/zylisp](https://zylisp.github.io/zylisp/doc/cmd/zylisp)

In short, once compiled, you may pass a flag for one of the supported REPL
modes (e.g., `-ast`), or use `zylisp` as a CLI tool (i.e., compiler), with the
`-cli` flag.

General package reference documentation is available here:

- [https://zylisp.github.io/zylisp/doc](https://zylisp.github.io/zylisp/doc)

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

ZYLISP version: 0.8.0
Build: master@1e52cac, 2019-04-07T04:16:06Z
REPL Mode: AST
Go version: go1.12.2

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

## History

In December of 2013, [@jcla1](https://github.com/jcla1) pushed his initial
work on [Gisp](https://github.com/jcla1/gisp) to Github. He hacked on it over
the course of the next six months. In December of 2014, it was
[then picked up](https://github.com/masukomi/gisp) by
[@masukomi](https://github.com/masukomi) who hacked on it for about a month.
Later, in 2017, [@rcarmo](https://github.com/rcarmo) forked the original and
merged in [@masukomi](https://github.com/masukomi)'s changes, merged a
(still-open) pull request from [@m90](https://github.com/m90), and made a few
updates himself.

In early 2019 [@oubiwann](https://github.com/oubiwann) had created a
[zylisp](https://github.com/zylisp) Github org where various Go Lisp's were
being explored. In particular, the
[zygomys](https://github.com/glycerine/zygomys) project (which was based on
a different Lisp/Go lineage, [Glisp](https://github.com/zhemao/glisp), which
was started in 2014). The ZYLISP Github org took inspiration in its name from
the zygomys project. However, work with zygomys was abandoned after it became
clear that Go interop was very awkward in this Lisp dialect. It was at this
point that [@rcarmo](https://github.com/rcarmo)'s fork was forked into the
ZYLISP org where it was eventually renamed from Gisp to ZYLISP.

### Versions

None of the previous forks tagged any of the work with versions, as such, this
fork has retroactrively tagged the various phases of the project's work with
the following:

| Version | Date       | Repo     | Notes                                              |
| ------- | ---------- | -------- | -------------------------------------------------- |
| 0.8.0   | 2019-04-06 | zylisp   | Code rename to zylisp, cleanup, Go modules, docker |
| 0.7.0   | 2019-03-12 | zylisp   | CLI support, improved compiling options, logging   |
| 0.6.0   | 2019-02-28 | zylisp   | Compatibility release, docs updates                |
| 0.5.0   | 2017-08-25 | rcarmo   | Merged PR from m90, minor fixes and tweaks         |
| 0.4.0   | 2014-12-20 | masukomi | Refactoring, tests, error handling                 |
| 0.3.0   | 2014-06-29 | jcla1    | Updates and a merged PR from kedebug               |
| 0.2.0   | 2014-02-12 | jcla1    | Follow-up work, control structures, etc.           |
| 0.1.0   | 2014-01-25 | jcla1    | Core original work                                 |

## Credits

- @jcla1 for the initial implementation
- @masukomi for adding a number of tests and checks
- @m90 for README fixes
- The ZYLISP project for new development

## License

MIT

<!-- Named page links below: /-->

[logo]: https://avatars2.githubusercontent.com/u/48034771?s=250
[logo-large]: https://avatars2.githubusercontent.com/u/48034771
[travis]: https://travis-ci.org/zylisp/zylisp
[travis-badge]: https://travis-ci.org/zylisp/zylisp.png?branch=master
[tag-badge]: https://img.shields.io/github/tag/zylisp/zylisp.svg
[tag]: https://github.com/zylisp/zylisp/tags
[go-v]: https://img.shields.io/badge/Go-1.12-blue.svg
