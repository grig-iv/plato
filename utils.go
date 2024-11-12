package main

import (
	"strings"

	"github.com/gobwas/glob"
)

func match(dir, pattern string) bool {
	pattern = replaceTilde(pattern)
	pattern = strings.ToLower(pattern)
	dir = strings.ToLower(dir)

	g := glob.MustCompile(pattern)
	return g.Match(dir)
}

func replaceTilde(dirPath string) string {
	if strings.HasPrefix(dirPath, "~/") {
		return strings.Replace(dirPath, "~/", homeDir, 1)
	}

	return dirPath
}
