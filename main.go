package main

import (
	"fmt"
	"log"
	"os"

	"github.com/laurence-trippen/git-ignit/github"
	"github.com/laurence-trippen/git-ignit/gitignit"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please specify a .gitignore File!")
		os.Exit(1)
	}

	gitignore := os.Args[1]
	gitignore = gitignit.CompleteIgnoreName(gitignore)

	contents, err := github.FetchRawRepoFile("github", "gitignore", "main", gitignore)
	if nil != err {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Println("Creating .gitignore ...")

	f, err := os.OpenFile(".gitignore", os.O_CREATE|os.O_EXCL, 0644)
	if nil != err {
		log.Fatal(err)

		// TODO: Check if closing in err handling is valid!? Might be nil...
		f.Close()

		os.Exit(1)
	}

	defer f.Close()

	f.WriteString(contents)

	fmt.Println("Done!")
}
