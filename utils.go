package main

import (
	"strings"

	"github.com/gobwas/glob"
)

func match(dir, pattern string) bool {
	g := glob.MustCompile(replaceTilde(pattern))
	return g.Match(dir)
}

func replaceTilde(dirPath string) string {
	if strings.HasPrefix(dirPath, "~/") {
		return strings.Replace(dirPath, "~/", homeDir, 1)
	}

	return dirPath
}
