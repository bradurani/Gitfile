# Gitfile

Installs git repos onto your system and keeps them up-to-date. It's a
lightweight package manager for things that haven't been published to a real
package manager. It's useful for installing and updating all the odd one-off things
that only live on [GitHub](https://github.com)

## Usage

List your packages in a [YAML](http://yaml.org) file called `Gitfile`
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

And your packages will be installed or updated

# Configuration

The `Gitfile` must be in YAML format and the top level element must be an array.
Options are

 - `url` - The Url (https or ssh) of the git repo
 - `path` - The path to install to
 - `branch` - The branch to install
 - `tag` - The tag to install

`tag` and `branch` cannot both be defined for any repo. If neither are defined,
then branch `master` is installed.

# Options

`gitfile` - Installs from `Gitfile` in the current directory
`gitfile -f <path>` - Installs using config file at the specified path. If path is
                      a dir, it will look for a `Gitfile`

# Installing

Gitfile is written in [Go](https://golang.org/). Package and instructions coming
soon.

# Contributing

Feel feel free to open issues and pull requests. If you like this repo, spread
the word!

### Potential Improvements
 - `--help` flag
 - `man` pages
 - `commit: ` config option
 - `post_install: ` config option for running script
 - `post_update: ` config option for running script
 - `gitfile status` command (show repo status)
 - `gitfile update` command (updates only, does not install new repos)
 - brew, deb, arch, yum etc. packages







