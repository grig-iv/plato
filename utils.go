package main

import (
	"os"
	"strings"

	"github.com/gobwas/glob"
)

var (
	homeDir       = os.Getenv("HOME") + "/"
	workingDir, _ = os.Getwd()
)

func match(pattern string) bool {
	g := glob.MustCompile(replaceTilde(pattern))
	return g.Match(workingDir)
}

func replaceTilde(dirPath string) string {
	if strings.HasPrefix(dirPath, "~/") {
		return strings.Replace(dirPath, "~/", homeDir, 1)
	}

	return dirPath
}
