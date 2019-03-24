package common

const (
	// Files and Directories Errors
	DirectoryError         string = "There was a problem accessing directory %s: %s"
	DirectoryCreationError string = "There was a problem creating directory %s: %s"
	// Implementation Errors
	NotImplementedError string = "Not implemented yet"

	// Unsupported Errors
	LogLevelUnsupportedError string = "The provided logging level is not supported"
)
