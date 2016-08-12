package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"net/url"
	"gopkg.in/yaml.v2"
	// "github.com/bradurani/go-git/git"
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
	updateRepos(repos)
}

func updateRepos(repos []repo) {
	for _, repo := range repos {
		updateRepo(repo)
	}
}

func updateRepo(repo repo) {
	path := repo.Path
	if path == "" {
		path = "."
	}
	fmt.Println("path is ", path)
	gitDir := parseGitDir(repo.Url)
	repoExists := repoExists(path, gitDir)
	fmt.Println("exists: ", repoExists)
}

func parseGitDir(repoUrl string) (gitDir string) {
	u, err := url.Parse(repoUrl)
	check(err)
	path := strings.TrimLeft(u.Path, "/")
	segments := strings.Split(path, "/")
	if len(segments) != 2 {
		panic("urls must have 2 path segments")
	}
	return segments[1]
}

func repoExists(repoPath string, gitDir string) (exists bool) {
	err := os.MkdirAll(repoPath, 0777)
	check(err)
	gitPath := filepath.Join(repoPath, gitDir, ".git")
	pathExists, err := pathExists(gitPath)
	check(err)
	return pathExists
}

func pathExists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return true, err
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
