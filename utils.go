package main

import (
	"log"
	"os"
	"path/filepath"
)

func runInChangedProjects(command string) string {
	return `bash "$(dirname -- "${BASH_SOURCE[0]}")/lib/run-in-changed-projects.sh" go.mod ` + command + ` -- "$@"`
}

func createDir(dir string) {
	wd, _ := os.Getwd()
	destPath := filepath.Join(wd, dir)

	err := os.MkdirAll(filepath.Join(destPath), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}
