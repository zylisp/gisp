package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/zylisp/zylisp/common"
	"github.com/zylisp/zylisp/repl"
)

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
