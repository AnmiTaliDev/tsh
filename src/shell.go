// Tiny Shell (tsh) - Command Executor
// Created: 2025-04-03 06:19:06 UTC
// Author: AnmiTaliDev

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Shell struct {
	reader    *bufio.Reader
	isRunning bool
}

func NewShell() *Shell {
	return &Shell{
		reader:    bufio.NewReader(os.Stdin),
		isRunning: true,
	}
}

func (s *Shell) Run() {
	for s.isRunning {
		fmt.Print(PROMPT)
		
		input, err := s.reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println()
				return
			}
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)
		
		switch input {
		case "":
			continue
		case "exit":
			s.isRunning = false
		case "help":
			printHelp()
		default:
			if err := s.executeCommand(input); err != nil {
				fmt.Fprintln(os.Stderr, "Error:", err)
			}
		}
	}
}

func (s *Shell) executeCommand(input string) error {
	args := strings.Fields(input)
	if len(args) == 0 {
		return nil
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}