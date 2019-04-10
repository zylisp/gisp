package repl

import (
	"bufio"
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"io"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
	"github.com/zylisp/zylisp/core/generator"
	"github.com/zylisp/zylisp/core/parser"
)

// ASTMain runs the main loop for the AST-based REPL
func ASTMain() {
	log.Info("Starting main loop ...")
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
		handleSignals()
		// line, _, _ := r.ReadLine()
		line, _, err := r.ReadLine()
		log.Debugf("Got: %#v, %#v", line, err)
		handleReadlineErrors(err)
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

// GoGenMain runs the main loop for the Go-code-based REPL
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

// LispMain runs the main loop for the Lisp-based REPL
func LispMain() {
	log.Info("Starting main loop ...")
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
		line, aa, bb := r.ReadLine()
		log.Tracef("Got: %#v, %#v, %#v", line, aa, bb)
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

func handleSignals() {
	log.Trace("Setting up signal handler ...")
	signalChan := make(chan os.Signal, 1)

	signal.Notify(signalChan,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		s := <-signalChan
		log.Debugf("Signal: %#v", s)
		switch s {
		case syscall.SIGINT: // ^C
			log.Debug("Received SIGNINT signal; quitting ...")
			fmt.Printf("\n%s\n", REPLCommonExitMsg)
			os.Exit(0)
		case syscall.SIGTERM:
			log.Debug("Received SIGTERM signal; quitting ...")
			os.Exit(0)
		case syscall.SIGQUIT:
			log.Debug("Received SIGQUIT signal; quitting ...")
			os.Exit(1)
		default:
			log.Debugf("Received unexpected signal %#v", s)
		}
	}()
}

func handleReadlineErrors(err error) {
	switch err {
	case io.EOF:
		log.Debug("Received EOF from input; quitting ...")
		fmt.Printf("\n%s\n", REPLCommonExitMsg)
		os.Exit(0)
	default:
		log.Debugf("Not able to handle error %#v", err)
	}
}
