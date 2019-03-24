package repl

import (
	"bufio"
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/zylisp/zylisp/generator"
	"github.com/zylisp/zylisp/parser"
)

func ASTMain() {
	banner := Banner{
		commonHelp: CommonREPLHelp,
		greeting:   REPLBannerGreeting,
		modeHelp:   ASTREPLHelp,
		replMode:   "AST",
	}

	banner.printBanner()
	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(ASTPrompt)
		line, _, _ := r.ReadLine()
		p := parser.ParseFromString("<REPL>", string(line)+"\n")
		log.Info("Parsed AST")
		log.Debugf("AST: %s", p)

		// a := generator.GenerateAST(p)
		a := generator.EvalExprs(p)
		fset := token.NewFileSet()
		ast.Print(fset, a)

		var buf bytes.Buffer
		printer.Fprint(&buf, fset, a)
		fmt.Printf("%s\n", buf.String())
	}
}

func GoGenMain() {
	banner := Banner{
		commonHelp: CommonREPLHelp,
		greeting:   REPLBannerGreeting,
		modeHelp:   GoGenREPLHelp,
		replMode:   "GOGEN",
	}

	banner.printBanner()
	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(GoGenPrompt)
		line, _, _ := r.ReadLine()
		log.Info("Generated Go code")
		generator.PrintGoFromLispString(string(line))
	}
}

func LispMain() {
	banner := Banner{
		commonHelp: CommonREPLHelp,
		greeting:   REPLBannerGreeting,
		modeHelp:   LispREPLHelp,
		replMode:   "Lisp",
	}

	banner.printBanner()
	r := bufio.NewReader(os.Stdin)

	// XXX we should explore REPL-based packages ... that would allow for a
	//     more Go-like experience in the REPL, with the ability to declare a
	//     new package in the REPL, and then refer to work done in the same
	//     session, but in a different REPL package ... I guess this applies
	//     more to the Lisp REPL
	for {
		fmt.Print(LispPrompt)
		line, _, _ := r.ReadLine()
		p := parser.ParseFromString("<REPL>", string(line)+"\n")
		log.Info("Parsed AST")
		log.Debugf("AST: %s", p)

		// a := generator.GenerateAST(p)
		a := generator.EvalExprs(p)
		fset := token.NewFileSet()
		ast.Print(fset, a)

		var buf bytes.Buffer
		printer.Fprint(&buf, fset, a)
		fmt.Printf("%s\n", buf.String())
	}
}
