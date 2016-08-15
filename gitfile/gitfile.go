package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

type repo struct {
	Url    string
	Path   string
	Branch string
	Tag    string
	Commit string
}

func main() {
	argsWithProg := os.Args
	gitfilePath := gitfilePath(argsWithProg)
	fmt.Println(gitfilePath)
	contents := readFile(gitfilePath)
	repos := parseFile(contents)
	addRepoDefaults(repos)
	fmt.Println(repos)
	updateRepos(repos)
}

func addRepoDefaults(repos []repo) {
	for i := range repos {
		if repos[i].Url == "" {
			panic("repos must have url field")
		}
		if repos[i].Path == "" {
			repos[i].Path = "."
		}
		if repos[i].Branch == "" && repos[i].Tag == "" && repos[i].Commit == "" {
			repos[i].Branch = "master"
		}
	}
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
	fmt.Println("path: ", path)
	gitDir := parseGitDir(repo.Url)
	repoExists := repoExists(path, gitDir)
	fmt.Println("exists: ", repoExists)
	if repoExists {
		pullRepo(repo, gitDir)
	} else {
		cloneRepo(repo, path)
	}
}

func pullRepo(repo repo, gitDir string) {
	pwd := getPwd()
	fmt.Println("pwd: ", pwd)
	changeDir(gitDir)
	args := []string{"pull", "--ff-only"}
	runCmd("/usr/local/bin/git", args)
	changeDir(pwd)
}

func changeDir(p string) {
	err := os.Chdir(p)
	check(err)
}

func getPwd() string {
	pwd, err := os.Getwd()
	check(err)
	return pwd
}

func cloneRepo(repo repo, path string) {
	args := []string{"clone", repo.Url}
	fmt.Println("path:", repo.Path)
	if strings.TrimSpace(repo.Path) != "" {
		args = append(args, repo.Path)
	}
	runCmd("/usr/local/bin/git", args)
}

func runCmd(cmd string, args []string) {
	fmt.Println("cmd: ", strings.Join(append([]string{cmd}, args...), " "))
	out, err := exec.Command(cmd, args...).Output()
	check(err)
	fmt.Println(out)
}

func parseGitDir(repoUrl string) (gitDir string) {
	u, err := url.Parse(repoUrl)
	check(err)
	path := strings.TrimLeft(u.Path, "/")
	segments := strings.Split(path, "/")
	if len(segments) != 2 {
		panic("urls must have 2 path segments")
	}
	lastSegment := strings.TrimRight(segments[1], ".git")
	return lastSegment
}

func repoExists(repoPath string, gitDir string) (exists bool) {
	err := os.MkdirAll(repoPath, 0777)
	check(err)
	gitPath := filepath.Join(repoPath, gitDir, ".git")
	fmt.Println("path exists? ", gitPath)
	pathExists, err := pathExists(gitPath)
	check(err)
	return pathExists
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
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

func gitfilePath(argsWithProg []string) string {
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
		fmt.Println(e)
		panic(e)
	}
}
