package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/zylisp/zylisp/core/generator"
	"github.com/zylisp/zylisp/repl"
)

func main() {
	modes, inputs, outputs := parseCLIOptions()

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
					PrepareOutputFile(outputs.outDir)
					outputs.files = append(outputs.files, outputs.outDir)
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

// Dispatch Functions

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
