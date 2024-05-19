package gitignit

import "strings"

func CompleteIgnoreName(name string) string {
	if strings.HasSuffix(name, ".gitignore") {
		return name
	}

	return name + ".gitignore"
}
