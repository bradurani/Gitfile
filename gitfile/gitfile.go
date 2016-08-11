package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
)

func main() {
	argsWithProg := os.Args
	absPath := AbsPath(argsWithProg)
	fmt.Println(absPath)
}

func AbsPath(argsWithProg []string) string {
	dirArg := "."
	if len(argsWithProg) > 1 {
		dirArg = argsWithProg[1]
	}
	currentDir := CurrentDir()
	absPath := path.Join(currentDir, dirArg, "Gitfile")
	return absPath
}

func CurrentDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
