package common

import (
	"bytes"
	"os"
	"path/filepath"
	"regexp"

	log "github.com/sirupsen/logrus"
)

var camelingRegex = regexp.MustCompile("[0-9A-Za-z]+")

// CamelCase takes a string and a boolean and returns a camel-cased version of
//           the provided string. If the provided boolean is true,
func CamelCase(src string, capit bool) string {
	byteSrc := []byte(src)
	chunks := camelingRegex.FindAll(byteSrc, -1)
	for idx, val := range chunks {
		if idx > 0 || capit {
			chunks[idx] = bytes.Title(val)
		}
	}
	return string(bytes.Join(chunks, nil))
}

// RemoveExtension returns just the bare name of a path without the final
//                 '.*' ending.
func RemoveExtension(path string) string {
	extension := filepath.Ext(path)
	return path[0 : len(path)-len(extension)]
}

// IsDir returns a boolean depending upon whether the provided path points
//       to a directory or not.
func IsDir(path string) bool {
	file, err := os.Stat(path)
	if err != nil {
		log.Debugf(DirectoryError, path, err.Error())
		return false
	}
	if file.Mode().IsDir() {
		return true
	}
	return false
}
