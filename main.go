package main

import (
	"fmt"
	"os"
	"strconv"
	"bufio"
	"strings"
)

func main() {
	cmd := os.Args[1]
	args := os.Args[2:]

	if cmd == "generate-secret" {
		executeGenerateSecret(args)
	} else if cmd == "encrypt" {
		executeEncrypt(args)
	} else if cmd == "decrypt" {
		executeDecrypt(args)
	} else {
		printError("Invalid command")
		os.Exit(1)
	}
}

func executeGenerateSecret(args []string) {
  lengthParam := args[0]

	length, err := strconv.Atoi(lengthParam)
	if err != nil {
		printError("Invalid length: %s", lengthParam)
		os.Exit(1)
	}

	secret, err := generateSecret(length)
	if err != nil {
		printError("Error generating secret: %s", err.Error())
		os.Exit(1)
	}

	printOutput(secret)
}

func executeEncrypt(args []string) {
  secret := args[0]
	initVector := args[1]

  var text string

  if len(args) > 2 {
    text = args[2]
  }

  if text == "" {
    text = readStdIn()
  }

	encrypted, err := encrypt(secret, initVector, text)
	if err != nil {
		printError("Error encrypting text: %s %s", text, err.Error())
		os.Exit(1)
	}

	printOutput(encrypted)
}

func executeDecrypt(args []string) {
  secret := args[0]
	initVector := args[1]

  var text string

  if len(args) > 2 {
    text = args[2]
  } 

  if text == "" {
    text = readStdIn()
  }

	decrypted, err := decrypt(secret, initVector, text)
	if err != nil {
		printError("Error decrypting text: %s %s", text, err.Error())
		os.Exit(1)
	}

	printOutput(decrypted)
}

func readStdIn() string {
	value := strings.Builder{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		value.WriteString(scanner.Text())
	}

	return value.String()
}

func printError(format string, args ...any) {
  fmt.Fprintf(os.Stderr, format + "\n", args...)
}

func printOutput(format string) {
  fmt.Fprint(os.Stdout, format)
}

