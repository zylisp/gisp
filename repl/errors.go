package repl

const (
	// Mode Errors
	ModeNeededError string = "You need to supply a mode"

	// Files and Directories Errors
	FilesNeededError          string = "You need to provide at least one file upon which to operate"
	DirNeededError            string = "You must define an output directory when processing multiple files"
	UnexpectedFileOrDirError  string = "Unexpected error with CLI and output file or directory"
	UnexpectedFilesOrDirError string = "Unexpected error with CLI and output files or directory"

	// Compile Errors
	CompileError            string = "Couldn't compile file: %s"
	CompileWithoutFileError string = "Cannot compile file(s) without file(s) defined"
	CompoileWithoutCLIError string = "Can only compile to bytecode in CLI mode"

	// Unsupported Errors
	LispCLIUnsupportedError  string = "The Lisp CLI is currently not supported"
	LispREPLUnsupportedError string = "Lisp mode is currently not supported"
)
