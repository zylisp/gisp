package repl

import (
	"fmt"
	"runtime"

	"github.com/zylisp/zylisp/common"
)

// Banner contains the data used for different banner types.
type Banner struct {
	commonHelp string
	greeting   string
	modeHelp   string
	replMode   string
}

func (b Banner) printGreeting() {
	fmt.Println(b.greeting)
}

func (b Banner) printHelp() {
	fmt.Print(b.commonHelp)
	fmt.Println(b.modeHelp)
}

func (b Banner) printVersions() {
	fmt.Printf("ZYLISP version: %s\n", common.VersionString())
	fmt.Printf("Build: %s\n", common.BuildString())
	fmt.Printf("REPL Mode: %s\n", b.replMode)
	fmt.Printf("Go version: %s\n", runtime.Version())
}

func (b Banner) printBanner() {
	b.printGreeting()
	b.printVersions()
	b.printHelp()
}
