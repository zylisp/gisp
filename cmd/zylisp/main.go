package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/zylisp/zylisp/common"
	"github.com/zylisp/zylisp/core/generator"
	"github.com/zylisp/zylisp/repl"
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
	useDir  bool
	useFile bool
}

// PrepareOutputDir ...
func PrepareOutputDir(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Infof("Directory '%s' does not exist; creating ...", dir)
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			msg := fmt.Sprintf(common.DirectoryError, dir, err.Error())
			log.Error(msg)
			panic(msg)
		}
	}
}

// PrepareOutputFile ...
func PrepareOutputFile(filename string) {
	log.Debug("Preparing output file:", filename)
	basename := filepath.Dir(filename)
	log.Debug("Got basename:", basename)
	PrepareOutputDir(basename)
}

// MakeOutputFilename ...
func MakeOutputFilename(prefix string, inputFile string, extension string) string {
	var template string
	if extension == "" {
		template = "%s%s%s%s"
	} else {
		template = "%s%s%s.%s"
	}
	return fmt.Sprintf(
		template,
		prefix,
		string(os.PathSeparator),
		filepath.Base(common.RemoveExtension(inputFile)),
		extension)
}

func compileGo(infile string, outfile string) {
	log.Infof("Compiling %s ...", outfile)
	cmd := exec.Command("go", "build", "-o", outfile, infile)
	cmd.Env = append(os.Environ(), "GO111MODULE=on")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Debugf("Build error: %s", output)
		log.Errorf(repl.CompileError, err.Error())
	} else {
		log.Tracef("Output: %s", output)
	}
}

func dispatchLisp(modes Modes) {
	if modes.cli {
		// LISP CLI
		fmt.Println(repl.LispCLIUnsupportedError)
	} else {
		// LISP REPL
		fmt.Println(repl.LispREPLUnsupportedError)
		// repl.LispMain()
	}
}

func dispatchAST(modes Modes, inputs Inputs, outputs Outputs) {
	if modes.cli {
		// AST CLI
		for i, inputFile := range inputs.files {
			log.Infof("Processing file '%s' for AST output '%s' ...",
				inputFile, outputs.files[i])
			log.Debug("Use file for output?", outputs.useFile)
			log.Debug("Use directory for output?", outputs.useDir)
			if outputs.useFile {
				generator.WriteASTFromFile(inputFile, outputs.files[i])
			} else {
				generator.PrintASTFromFile(inputFile)
			}
		}
	} else {
		// AST REPL
		repl.ASTMain()
	}
}

func dispatchGoGen(modes Modes, inputs Inputs, outputs Outputs) {
	if modes.cli {
		// Go-generator CLI
		for i, inputFile := range inputs.files {
			log.Infof("Processing file '%s' for Go output '%s' ...",
				inputFile, outputs.files[i])
			if outputs.useFile {
				generator.WriteGoFromFile(inputFile, outputs.files[i])
			} else {
				generator.PrintGoFromFile(inputFile)
			}
		}
	} else {
		// GOGEN REPL
		repl.GoGenMain()
	}
}

func dispatchByteCode(modes Modes, inputs Inputs, outputs Outputs) {
	if modes.cli {
		// Go-compiler CLI
		for i, inputFile := range inputs.files {
			outputFile := outputs.files[i]
			goOutputFile := outputFile + ".go"
			log.Infof("Processing file '%s' for Go output '%s' ...",
				inputFile, goOutputFile)
			if outputs.useFile {
				generator.WriteGoFromFile(inputFile, goOutputFile)
			} else {
				log.Error(repl.CompileWithoutFileError)
			}
			log.Infof("Processing file '%s' for byte-code output '%s' ...",
				goOutputFile, outputFile)
			compileGo(goOutputFile, outputFile)
		}
	} else {
		log.Error(repl.CompoileWithoutCLIError)
	}
}

// XXX Currently unused; remove?
// func getUseDir (dir bool) bool {
// 	if dir {
// 		return true
// 	} else {
// 		return false
// 	}
// }

func getHasFiles(files []string) bool {
	if len(files) > 0 {
		return true
	}
	return false
}

func extensionFromMode(modes Modes) string {
	var extension string
	if modes.lisp {
		extension = "zsp"
	} else if modes.ast {
		extension = "ast"
	} else if modes.gogen {
		extension = "go"
	} else if modes.bytecode {
		extension = ""
	} else {
		fmt.Println(repl.ModeNeededError)
		os.Exit(1)
	}
	return extension
}

func dispatch(modes Modes, inputs Inputs, outputs Outputs) {
	log.Debug("Dispatched")
	log.Debug("Got modes:", modes)
	log.Debug("Got inputs:", inputs)
	log.Debug("Got outputs:", outputs)
	if modes.lisp {
		dispatchLisp(modes)
	} else if modes.ast {
		dispatchAST(modes, inputs, outputs)
	} else if modes.gogen {
		dispatchGoGen(modes, inputs, outputs)
	} else if modes.bytecode {
		dispatchByteCode(modes, inputs, outputs)
	} else {
		fmt.Println(repl.ModeNeededError)
		os.Exit(1)
	}
}

func main() {
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
		useDir:  isDir,
		useFile: false,
	}

	if modes.cli {
		// Check for at least one file to operate upon, when in CLI mode
		if inputs.hasFiles {
			if len(inputs.files) > 1 {
				log.Debug("Got multiple input files")
				// If more than one file is given, ignore output file and only use dir
				if outputs.isDir {
					log.Debug("Output dir is defined: using it ...")
					outputs.files = []string{}
					// Since we're going to be using the dir, make sure it exists/create
					// if necessary
					PrepareOutputDir(outputs.dir)
					log.Debug("Input files:", inputs.files)
					for _, file := range inputs.files {
						outputs.files = append(outputs.files, MakeOutputFilename(
							outputs.dir, file, extensionFromMode(modes)))
					}
					outputs.useFile = true
				} else {
					log.Error(repl.DirNeededError)
					os.Exit(1)
				}
			} else {
				log.Debug("Got a single input file")
				log.Debug("Original output files:", outputs.files)
				log.Debug("Original input files:", inputs.files)
				// if only one file is given and dir is given, then set the output file to
				// be the dir/infile.updated-extension
				if outputs.isDir {
					log.Debug("Outputs is a directory, using it ...")
					PrepareOutputDir(outputs.dir)
					outputs.files = append(outputs.files, MakeOutputFilename(
						outputs.dir, inputs.files[0], extensionFromMode(modes)))
					outputs.useFile = true
					log.Debug("Modified output files:", outputs.files)
					log.Debug("Modified input files:", inputs.files)
					// If only one file is given, and the output file is set
				} else {
					log.Debug("Outputs is not a directory.")
					PrepareOutputFile(*outPtr)
					outputs.files = append(outputs.files, *outPtr)
					outputs.useFile = true
					log.Debug("Modified output files:", outputs.files)
					log.Debug("Modified input files:", inputs.files)
				}
			}
		} else {
			log.Error(repl.FilesNeededError)
			os.Exit(1)
		}
	}
	log.Debug("Preparing to dispatch ...")
	dispatch(modes, inputs, outputs)
}
