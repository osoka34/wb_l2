package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("Welcome to MyShell!")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("myshell> ")
		scanner.Scan()
		input := scanner.Text()

		if input == "exit" {
			break
		}
		if input == "" || input == "\n" {
			continue
		}

		commands := strings.Split(input, "|")

		var inputReader io.Reader = os.Stdin

		for _, cmd := range commands {
			r, w := io.Pipe()
			go func(command string, inputReader io.Reader, pipeWriter *io.PipeWriter) {
				defer w.Close()
				runCommand(command, inputReader, pipeWriter)
			}(cmd, inputReader, w)

			inputReader = r
		}

		io.Copy(os.Stdout, inputReader)
	}
	fmt.Println("Goodbye!")
}

func runCommand(command string, inputReader io.Reader, pipeWriter *io.PipeWriter) {
	command = strings.TrimSpace(command)
	parts := strings.Fields(command)
	//fmt.Println(parts)
	cmd := exec.Command(parts[0], parts[1:]...)
	cmd.Stdin = inputReader
	cmd.Stdout = pipeWriter
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
