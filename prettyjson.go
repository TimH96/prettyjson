/*
prettyjson.go
author: TimH96
*/

// Shell tool to pretty print json.
package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/akamensky/argparse"
)

// Type for command line arguments for script.
type CLArgs struct {
	Order  string
	Depth  int
	Indent int
}

// Read piped stdin and return resulting string.
// Inspired by: https://flaviocopes.com/go-shell-pipes/
func readStdin() (out string, err error) {
	info, err := os.Stdin.Stat()
	// throw piped error
	if err != nil {
		panic(err)
	}
	// panic error when no piped stdin
	if (info.Mode() & os.ModeNamedPipe) == 0 {
		return "", errors.New("No piped input for prettyjson, use prettyjson --help for usage information")
	}
	// parse stdin to string input
	reader := bufio.NewReader(os.Stdin)
	var raw []rune
	for {
		char, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			err = nil
			break
		}
		raw = append(raw, char)
	}
	return string(raw), nil
}

// Parses command line arguments and returns them.
func getArgs() (args CLArgs, terminal bool, err error) {
	// define parse
	parser := argparse.NewParser("prettyjson", "Pretty prints provided json string to stdout")
	order := parser.String("o", "order", &argparse.Options{Required: false, Help: "Key order", Default: nil})
	depth := parser.Int("d", "depth", &argparse.Options{Required: false, Help: "Recursion depth", Default: -1})
	indent := parser.Int("i", "indent", &argparse.Options{Required: false, Help: "Indent per level", Default: 4})
	// parse input and return resulting struct
	err = parser.Parse(os.Args)
	return CLArgs{
		Order:  *order,
		Depth:  *depth,
		Indent: *indent,
	}, false, err
}

// Script entrypoint.
func main() {
	// get args
	args, _, _ := getArgs()
	fmt.Print(args)
	// get piped input
	input, err := readStdin()
	if err != nil {
		fmt.Print(err)
		return
	}
	// parse dynamic json map
	var result map[string]interface{}
	err = json.Unmarshal([]byte(input), &result)
	if err != nil {
		fmt.Print(err)
		return
	}
	// cont
	fmt.Print(result)
}
