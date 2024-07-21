package main

import (
	"fmt"
	"os"
	"strconv"
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
	params, err := ParseArgs(args, []Parameter{
		Param("length", true, "32"),
	})

	if err != nil {
		printError(err.Error())
		os.Exit(1)
	}

	length, err := strconv.Atoi(params.Get("length"))
	if err != nil {
		printError("Invalid length: %s", params.Get("length"))
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
	params, err := ParseArgs(args, []Parameter{
		Param("secret", true, ""),
		Param("iv", true, ""),
		StdInParam("text", true),
	})

	if err != nil {
		printError(err.Error())
		os.Exit(1)
	}

	text := params.Get("text")
	secret := params.Get("secret")
	initVector := params.Get("iv")

	encrypted, err := encrypt(secret, initVector, text)
	if err != nil {
		printError("Error encrypting text: %s %s", text, err.Error())
		os.Exit(1)
	}

	printOutput(encrypted)
}

func executeDecrypt(args []string) {
	params, err := ParseArgs(args, []Parameter{
		Param("secret", true, ""),
		Param("iv", true, ""),
		StdInParam("text", true),
	})

	if err != nil {
		printError(err.Error())
		os.Exit(1)
	}

	text := params.Get("text")
	secret := params.Get("secret")
	initVector := params.Get("iv")

	decrypted, err := decrypt(secret, initVector, text)
	if err != nil {
		printError("Error decrypting text: %s %s", text, err.Error())
		os.Exit(1)
	}

	printOutput(decrypted)
}

func printError(format string, args ...any) {
  fmt.Fprintf(os.Stderr, format + "\n", args...)
}

func printOutput(format string, args ...any) {
  fmt.Fprintf(os.Stdout, format, args...)
}

