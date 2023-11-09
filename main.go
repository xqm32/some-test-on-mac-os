package main

import (
	"os"
	"path/filepath"
)

func main() {
	exeOld, err := os.Executable()
	println("exeOld is ", exeOld)
	if err != nil {
		panic(err)
	}

	tempDir := os.TempDir()
	exeNewDir := filepath.Join(tempDir, "path/to/exeNew")
	if err := os.MkdirAll(exeNewDir, 0755); err != nil {
		panic(err)
	}

	exeNew := filepath.Join(exeNewDir)
	println("exeNew is ", exeNew)
	exeOldRel, err := filepath.Rel(exeNewDir, exeOld)
	println("exeOldRel is ", exeOldRel)
	if err != nil {
		panic(err)
	}

	if err := os.Symlink(exeOldRel, exeNew); err != nil {
		panic(err)
	}

	p, err := filepath.EvalSymlinks(exeNew)
	if err != nil {
		panic(err)
	}

	println(p)
}
