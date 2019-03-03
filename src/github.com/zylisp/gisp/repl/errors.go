package repl

// Mode Errors
const ModeNeededError string = "You need to supply a mode"

// Files and Directories Errors
const FilesNeededError string = "You need to provide at least one file upon which to operate"
const DirNeededError string = "You must define an output directory when processing multiple files"
const UnexpectedFileOrDirError string = "Unexpected error with CLI and output file or directory"
const UnexpectedFilesOrDirError string = "Unexpected error with CLI and output files or directory"

// Unsupported Errors
const LispCLIUnsupportedError string = "The Lisp CLI is currently not supported"
const LispREPLUnsupportedError string = "Lisp mode is currently not supported"
