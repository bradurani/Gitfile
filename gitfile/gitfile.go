package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"os/exec"
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
	pwd := getPwd()
	changeDir(gitfilePath)
	contents := readFile(filepath.Join(gitfilePath, "Gitfile"))
	repos := parseFile(contents)
	addRepoDefaults(repos)
	updateRepos(repos)
	changeDir(pwd)
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
	repoDir := parseRepoDir(repo.Url)
	repoExists := repoExists(repo.Path, repoDir)
	if repoExists {
		fetchRepo(repo, repoDir)
		checkout(repo, repoDir)
	} else {
		cloneRepo(repo, repoDir)
		checkout(repo, repoDir)
	}
}

func checkout(repo repo, repoDir string) {
	fullPath := filepath.Join(repo.Path, repoDir)
	pwd := getPwd()
	changeDir(fullPath)
	if repo.Commit != "" {
		runGitCmd([]string{"checkout", repo.Commit})
	} else if repo.Tag != "" {
		tagArg := fmt.Sprintf("tags/%s", repo.Tag)
		runGitCmd([]string{"checkout", tagArg})
	} else if repo.Branch != "" {
		runGitCmd([]string{"checkout", repo.Branch})
		runGitCmd([]string{"pull", "--ff-only"})
	} else {
		panic("No checkout value")
	}
	changeDir(pwd)
}

func fetchRepo(repo repo, repoDir string) {
	fullPath := filepath.Join(repo.Path, repoDir)
	fmt.Println("fetching repo", repo)
	pwd := getPwd()
	changeDir(fullPath)
	args := []string{"fetch"}
	runGitCmd(args)
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

func cloneRepo(repo repo, repoDir string) {
	fmt.Println("Cloning repo", repo)
	args := []string{"clone", repo.Url}
	if strings.TrimSpace(repo.Path) != "." {
		args = append(args, filepath.Join(repo.Path, repoDir))
	}
	runGitCmd(args)
}

func runGitCmd(args []string) {
	runCmd("git", args)
}

func runCmd(cmd string, args []string) {
	fullCmd := strings.Join(append([]string{cmd}, args...), " ")
	fmt.Println(fullCmd)
	out, err := exec.Command("sh", "-c", "'" + fullCmd + "'").Output()
	check(err)
	fmt.Printf("%s", out)
}

func parseRepoDir(repoUrl string) (gitDir string) {
	u, err := url.Parse(repoUrl)
	check(err)
	path := strings.TrimLeft(u.Path, "/")
	segments := strings.Split(path, "/")
	if len(segments) != 2 {
		panic("urls must have 2 path segments")
	}
	lastSegment := strings.TrimSuffix(segments[1], ".git")
	return lastSegment
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
	absPath, err := filepath.Abs(dirArg)
	check(err)
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
