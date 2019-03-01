package repl

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/zylisp/gisp"
	"github.com/zylisp/gisp/generator"
	"github.com/zylisp/gisp/parser"
	"go/ast"
	"go/printer"
	"go/token"
	"os"
	"runtime"
)

func AstMain() {
	r := bufio.NewReader(os.Stdin)

  fmt.Println(ReplBanner)
  fmt.Printf("ZYLISP version: %s [AST mode]\n", gisp.Version())
  fmt.Printf("Go version: %s\n", runtime.Version())
  fmt.Print(CommonReplHelp)
  fmt.Println(AstReplHelp)

	for {
		fmt.Print(AstPrompt)
		line, _, _ := r.ReadLine()
		p := parser.ParseFromString("<REPL>", string(line)+"\n")
		fmt.Printf("Parsed:\n%s\n", p)

		// a := generator.GenerateAST(p)
		a := generator.EvalExprs(p)
		fset := token.NewFileSet()
		fmt.Println("AST:")
		ast.Print(fset, a)

		var buf bytes.Buffer
		printer.Fprint(&buf, fset, a)
		fmt.Printf("%s\n", buf.String())
	}
}
