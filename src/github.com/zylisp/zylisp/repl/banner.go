package repl

import (
	"fmt"
	"github.com/zylisp/zylisp"
	"runtime"
)

type Banner struct {
	commonHelp string
	greeting string
	modeHelp string
	replMode string
}

func (b Banner) printGreeting() {
	fmt.Println(b.greeting)
}

func (b Banner) printHelp() {
  fmt.Print(b.commonHelp)
  fmt.Println(b.modeHelp)
}

func (b Banner) printVersions() {
	fmt.Printf("ZYLISP version: %s\n", zylisp.VersionString())
	fmt.Printf("Build: %s\n", zylisp.BuildString())
  fmt.Printf("REPL Mode: %s\n", b.replMode)
  fmt.Printf("Go version: %s\n", runtime.Version())
}

func (b Banner) printBanner() {
	b.printGreeting()
	b.printVersions()
	b.printHelp()
}
