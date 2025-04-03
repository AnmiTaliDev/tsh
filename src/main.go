// Tiny Shell (tsh)
// Created: 2025-04-03 06:19:06 UTC
// Author: AnmiTaliDev

package main

import (
	"fmt"
	"os"
)

const (
	VERSION       = "1.0.0"
	PROMPT        = "tsh> "
	SHELL_NAME    = "Tiny Shell"
	AUTHOR        = "AnmiTaliDev"
	BUILD_DATE    = "2025-04-03 06:19:06"
	EXIT_SUCCESS  = 0
	EXIT_ERROR    = 1
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-v", "--version":
			fmt.Printf("%s v%s (%s)\n", SHELL_NAME, VERSION, BUILD_DATE)
			return
		case "-h", "--help":
			printHelp()
			return
		default:
			fmt.Fprintf(os.Stderr, "Unknown option: %s\n", os.Args[1])
			fmt.Fprintf(os.Stderr, "Use --help for usage information\n")
			os.Exit(EXIT_ERROR)
		}
	}

	shell := NewShell()
	shell.Run()
}

func printHelp() {
	fmt.Printf(`%s v%s

Usage: tsh [OPTIONS]

Options:
  -h, --help     Show this help message
  -v, --version  Show version information

Built-in commands:
  exit           Exit the shell
  help           Show this help message

Author: %s
Build Date: %s
`, SHELL_NAME, VERSION, AUTHOR, BUILD_DATE)
}