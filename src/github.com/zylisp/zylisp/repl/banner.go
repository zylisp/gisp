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
	fmt.Printf("ZYLISP version: %s [%s mode]\n", zylisp.Version(), b.replMode)
  fmt.Printf("Go version: %s\n", runtime.Version())
}

func (b Banner) printBanner() {
	b.printGreeting()
	b.printVersions()
	b.printHelp()
}
