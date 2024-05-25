# Git Ignit

Download Github .gitnore templates superfassst!

![screenshot](docs/img/gitignit-cli-win.png)

### Applied Knowledge

- [x] Building CLIs in Go
- [x] GitHub Actions for CI/CD
- [x] Go Cross-Compilation
- [x] Simple Unit-Testing

### Installation

Coming soon...

### Building

```bash
$ go build
$ go install
```


### Testing

In the projects root directory run:

`$ go test -v ./...`

### Usage

Shorthand:
`$ git-ignit Java`

Full-Name:
`$ git-ignit Swift.gitignore`

### Behaviour

* The `.gitignore` file will be created in the current working / caller directory.
* Already existing .gitignore(s) won't be overwritten.

### Features

- [ ] GitHub Workflow CI/CD
  - [x] Build & Test Workflows on Pull Request
  - [x] Go Cross-Compile Artifacts Workflow on manual dispatch
  - [ ] Release workflow with SemVer & deployment for...
    - [ ] Windows MSI & ZIP
    - [ ] Windows WinGet?
    - [ ] Windows Choco?
    - [ ] Linux Debian Apt
    - [ ] Linux Flatpak
    - [ ] macOS HomeBrew
* [ ] File Overwrite optional option
* [ ] Cache template files
* [ ] Add & manage custom templates
* [X] Unit-Tests
