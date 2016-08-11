package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type repo struct {
	Url    string
	Path   string
	Branch string
	Tag    string
}

func main() {
	argsWithProg := os.Args
	absPath := absPath(argsWithProg)
	fmt.Println(absPath)
	contents := readFile(absPath)
	repos := parseFile(contents)
	fmt.Println(repos)
}

func readFile(absPath string) string {
	dat, err := ioutil.ReadFile(absPath)
	check(err)
	return string(dat)
}

func parseFile(contents string) []repo {
	r := []repo{}
	err := yaml.Unmarshal([]byte(contents), &r)
	check(err)
	fmt.Printf("--- t:\n%v\n\n", r)
	return r
}

func absPath(argsWithProg []string) string {
	dirArg := "."
	if len(argsWithProg) > 1 {
		dirArg = argsWithProg[1]
	}
	currentDir := currentDir()
	absPath := path.Join(currentDir, dirArg, "Gitfile")
	return absPath
}

func currentDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	check(err)
	return dir
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
