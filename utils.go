package main

import (
	"log"
	"os"
	"path/filepath"
)

func createDir(dir string) {
	wd, _ := os.Getwd()
	destPath := filepath.Join(wd, dir)

	err := os.MkdirAll(filepath.Join(destPath), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}
