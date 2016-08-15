# Gitfile

Installs git repos onto your system and keeps them up-to-date. It's a
lightweight package manager for things that haven't been published to a real
package manager. It's useful for installing and updating all the odd one-off things
that only live on [GitHub](https://github.com)

## Usage

List the repos you want installed in a [YAML](http://yaml.org) file called `Gitfile`
```yaml
# ~/my/source/dir/Gitfile

- url: git://github.com/bradurani/bradrc.git
- url: https://github.com/thoughtbot/dotfiles.git
  path: thoughtbot/
- url: https://github.com/olivierverdier/zsh-git-prompt.git
  tag: v0.4
- url: https://github.com/tmux-plugins/tmux-battery.git
  path: tmux-plugins/
  branch: master
```

Run `gitfile`
```bash
cd ~/my/source/dir
gitfile
```

And your repos will be cloned or fetched

# Configuration

The `Gitfile` must be in YAML format and the top level element must be an array.
Options are

 - `url` - The Url (https or ssh) of the git repo
 - `path` - The path to install to (absolute or relative to current dir)
 - `branch` - The branch to install
 - `tag` - The tag to install
 - `commit` - The commit to install

You can only use one of `tag`, `branch`, and `repo`. If none are defined,
then branch `master` is installed.

# Options

`gitfile` - Installs from `Gitfile` in the current directory  
`gitfile <dir>` - Installs using the `Gitfile` in the specified directory. Repos are
                  installed relative to the specified directory, not the directory the 
                  command is run from

# Installing

If you don't have Go, you must install it from [Go](https://golang.org/).
Then run:
```
go get github.com/bradurani/Gitfile/gitfile
go install github.com/bradurani/Gitfile/gitfile
```
# Contributing

Feel feel free to open issues and pull requests. If you like this repo, spread
the word!

### Potential Improvements
 - `--help` flag
 - `man` pages
 - `post_install: ` config option for running script
 - `post_update: ` config option for running script
 - `gitfile status` command (show repo status)
 - brew, deb, arch, yum etc. packages







