package main

import (
	"bytes"
	"fmt"
	"github.com/zylisp/gisp/generator"
	"github.com/zylisp/gisp/parser"
	"github.com/zylisp/gisp/repl"
	"go/printer"
	"go/token"
	"io/ioutil"
	"os"
)

func args(filename string) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	p := parser.ParseFromString(filename, string(b)+"\n")
	a := generator.GenerateAST(p)
	fset := token.NewFileSet()

	var buf bytes.Buffer
	printer.Fprint(&buf, fset, a)
	fmt.Printf("%s\n", buf.String())
}

func main() {
	if len(os.Args) > 1 {
		args(os.Args[1])
		return
	}
	repl.AstMain()
}
