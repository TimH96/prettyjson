/*
prettyjson.go
author: TimH96
*/

// shell tool to pretty print .json
package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

// type for command line arguments for script
type Args struct {
	order  string
	depth  int
	indent int
}

// read piped stdin and return resulting string
// inspired by https://flaviocopes.com/go-shell-pipes/
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

// parses command line arguments and returns them
func getArgs() (a Args, terminal bool) {
	return Args{
		"asc"
		5
		5
	}, true
}

// script entrypoint
func main() {
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
