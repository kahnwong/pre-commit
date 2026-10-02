package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// runInChangedProjects wraps a trusted shell command. The optional project marker
// defaults to go.mod and may be a filename glob, such as *.tf.
func runInChangedProjects(command string, markers ...string) string {
	marker := "go.mod"
	if len(markers) > 1 {
		panic("runInChangedProjects accepts at most one project marker")
	}
	if len(markers) == 1 {
		marker = markers[0]
	}
	quote := func(s string) string {
		return "'" + strings.ReplaceAll(s, "'", "'\"'\"'") + "'"
	}
	return `bash "$(dirname -- "${BASH_SOURCE[0]}")/lib/run-in-changed-projects.sh" ` + quote(marker) + ` bash -c ` + quote(command) + ` -- "$@"`
}

func createDir(dir string) {
	wd, _ := os.Getwd()
	destPath := filepath.Join(wd, dir)

	err := os.MkdirAll(filepath.Join(destPath), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}
