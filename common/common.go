package common

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func Oops(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}

func AbsPath(path string) string {
	wd, err := os.Getwd()
	Oops(err)
	return filepath.Join(wd, path)
}

func Lines(bytes []byte) []string {
	return strings.Split(string(bytes), "\n")
}
