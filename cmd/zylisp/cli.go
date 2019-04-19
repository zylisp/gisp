package main

import (
	"flag"
	"os"

	"github.com/zylisp/zylisp/common"
)

// Modes store how the zylisp executable has been run
type Modes struct {
	cli      bool
	ast      bool
	gogen    bool
	bytecode bool
	lisp     bool
}

// Inputs stores options about any files zylisp is operating upon
type Inputs struct {
	files    []string
	hasFiles bool
}

// Outputs stores options about the manner in which zylisp writes to files
type Outputs struct {
	dir     string
	files   []string
	isDir   bool
	isFile  bool
	outDir  string
	useDir  bool
	useFile bool
}

func parseCLIOptions() (Modes, Inputs, Outputs) {
	astPtr := flag.Bool("ast", false, "Enable AST mode")
	cliPtr := flag.Bool("cli", false, "Run as a CLI tool")
	dirPtr := flag.String("dir", "", "Default directory for writing operations")
	goPtr := flag.Bool("go", false, "Enable Go code-generation mode")
	byteCodePtr := flag.Bool("bytecode", false, "Enable byte-code compilation from generated Go")
	lispPtr := flag.Bool("lisp", false, "Enable LISP mode")
	logLevelPtr := flag.String("loglevel", "warning", "Set the logging level")
	outPtr := flag.String("o", "", "Default filename for writing operations")
	versionPtr := flag.Bool("version", false, "Display version/build info and exit")

	flag.Parse()
	if *versionPtr {
		println("Version: ", common.VersionString())
		println("Build: ", common.BuildString())
		os.Exit(0)
	}

	common.SetupLogger(*logLevelPtr)
	inputFiles := flag.Args()
	hasFiles := getHasFiles(inputFiles)
	isDir := len(*dirPtr) > 0

	modes := Modes{
		cli:      *cliPtr,
		ast:      *astPtr,
		gogen:    *goPtr,
		bytecode: *byteCodePtr,
		lisp:     *lispPtr,
	}

	inputs := Inputs{
		files:    inputFiles,
		hasFiles: hasFiles,
	}

	outputs := Outputs{
		dir:     *dirPtr,
		isDir:   isDir,
		isFile:  len(*outPtr) > 0,
		outDir:  *outPtr,
		useDir:  isDir,
		useFile: false,
	}
	return modes, inputs, outputs
}
