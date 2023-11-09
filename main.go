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

	exeNewDir := os.TempDir()
	exeNew := filepath.Join(exeNewDir, "exeNew")
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
