package github

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"slices"
	"strings"

	"github.com/gocolly/colly"
)

const kGithubTemplateUrl = "https://github.com/%s/%s"
const kGithubRawFileUrl = "https://raw.githubusercontent.com/%s/%s/%s/%s"

func FetchProjectPage(username, project string) {
	url := fmt.Sprintf(kGithubTemplateUrl, username, project)

	res, err := http.Get(url)
	if nil != err {
		fmt.Println("Failed to get GitHub Project!")
		return
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		return
	}

	fmt.Printf("client: response body: %s\n", resBody)
}

func FetchGitignoreFileNames() {
	c := colly.NewCollector()

	ignores := make([]string, 10)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)
	})

	c.OnHTML(".react-directory-filename-column .Link--primary", func(e *colly.HTMLElement) {
		repoFileName := e.Text

		if !strings.HasSuffix(repoFileName, ".gitignore") {
			return
		}

		if slices.Contains(ignores, repoFileName) {
			return
		}

		ignores = append(ignores, repoFileName)

		fmt.Println("Found: ", repoFileName)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, " scraped!")
	})

	c.Visit(fmt.Sprintf(kGithubTemplateUrl, "github", "gitignore"))
}

func FetchRawRepoFile(user, project, branch, path string) (string, error) {
	url := fmt.Sprintf(kGithubRawFileUrl, user, project, branch, path)

	res, err := http.Get(url)
	if nil != err {
		return "", errors.New("github: failed to fetch file")
	}

	resBody, err := io.ReadAll(res.Body)
	if nil != err {
		return "", errors.New("github: failed to read response body")
	}

	if res.StatusCode == 404 {
		return "", errors.New("github: gitignore not found")
	}

	fmt.Println("Found: " + url)

	return string(resBody), nil
}
