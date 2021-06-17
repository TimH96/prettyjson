/*
prettyjson.go
author: TimH96
*/

// shell tool to pretty print .json
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// package entrypoing
// much of the basic IO code was inspired by https://flaviocopes.com/go-shell-pipes/
func main() {
	info, err := os.Stdin.Stat()
	// throw piped error
	if err != nil {
		panic(err)
	}
	// panic error when no piped stdin
	if (info.Mode() & os.ModeNamedPipe) == 0 {
		fmt.Println("ERROR: No piped input for prettyjson, use prettyjson --help for usage infomration")
		return
	}
	// parse stdin to string input
	reader := bufio.NewReader(os.Stdin)
	var raw []rune
	for {
		char, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		raw = append(raw, char)
	}
	json_string := string(raw)
	// parse utf8 str
	fmt.Println(json_string)
}
