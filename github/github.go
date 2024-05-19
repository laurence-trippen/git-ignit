package github

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

const kGithubRawFileUrl = "https://raw.githubusercontent.com/%s/%s/%s/%s"

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
