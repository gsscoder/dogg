package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const (
	ProgramName    = "dogg"
	ProgramTitle   = "Polls processes for resources consumption and existence"
	ProgramVersion = "v0.4.0"
)

var (
	ConfigPath = fmt.Sprintf("./%s.yaml", ProgramName)
)

func Exit(message string, help bool) {
	fmt.Printf("%s: %s\n", ProgramName, message)
	if help {
		fmt.Printf("Try %s --help\n", ProgramName)
	}
	os.Exit(1)
}

func PrintInfo() {
	fmt.Printf("%s: %s\n", ProgramName, ProgramTitle)
	fmt.Printf("Version: %s\n", ProgramVersion)
}

func ReadText(path string) string {
	file, err := os.Open(path)
	if err == nil {
		content, err := ioutil.ReadAll(file)
		if err == nil {
			return string(content)
		}
	}
	return ""
}
