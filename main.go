package main

import (
	"os"
	"path/filepath"
)

func main() {
	exeOld, err := os.Executable()
	if err != nil {
		panic(err)
	}
	println("exeOld is", exeOld)

	exeOld, err = filepath.EvalSymlinks(exeOld)
	if err != nil {
		panic(err)
	}
	println("Real exeOld is", exeOld)

	tempDir := os.TempDir()
	println("tempDir is", tempDir)

	tempDir, err = filepath.EvalSymlinks(exeOld)
	if err != nil {
		panic(err)
	}
	println("Real tempDir is", tempDir)

	exeNewDir := filepath.Join(tempDir, "path/to")
	if err := os.MkdirAll(exeNewDir, 0755); err != nil {
		panic(err)
	}

	exeNew := filepath.Join(exeNewDir, "exeNew")
	println("exeNew is", exeNew)
	exeOldRel, err := filepath.Rel(exeNewDir, exeOld)
	println("exeOldRel is", exeOldRel)
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
