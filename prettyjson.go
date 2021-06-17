/*
prettyjson.go
author: TimH96
*/

// Shell tool to pretty print json.
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/akamensky/argparse"
)

const VERSION = "1.0"

// Command line arguments for script
type CLArgs struct {
	Indent  int
	Version bool
}

// Read piped stdin and return resulting string.
// Inspired by: https://flaviocopes.com/go-shell-pipes/
func readStdin() (out []byte, err error) {
	info, err := os.Stdin.Stat()
	// throw piped error
	if err != nil {
		panic(err)
	}
	// return error when no piped stdin
	if (info.Mode() & os.ModeNamedPipe) == 0 {
		return []byte{}, errors.New("No piped input for prettyjson, use prettyjson --help for usage information")
	}
	// read bytestream
	reader := bufio.NewReader(os.Stdin)
	raw, _, err := reader.ReadLine()
	return raw, err
}

// Parses command line arguments and returns them.
func getArgs() (args CLArgs, err error) {
	// define parse
	parser := argparse.NewParser("prettyjson", "Pretty prints provided json string to stdout")
	indent := parser.Int("i", "indent", &argparse.Options{Required: false, Help: "Indent per level", Default: 4})
	version := parser.Flag("v", "version", &argparse.Options{Required: false, Help: "Get script version"})
	err = parser.Parse(os.Args)
	// return resulting struct
	return CLArgs{
		Indent:  *indent,
		Version: *version,
	}, err
}

// Script entrypoint.
func main() {
	// get args
	args, err := getArgs()
	if err != nil {
		fmt.Print(err)
		return
	}
	if args.Version == true {
		fmt.Print("prettyjson version " + VERSION)
		return
	}
	// get piped input
	input, err := readStdin()
	if err != nil {
		fmt.Print(err)
		return
	}
	// prettify json string
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, input, "", strings.Repeat(" ", args.Indent))
	if err != nil {
		fmt.Print(err)
		return
	}
	// print result
	fmt.Print(string(prettyJSON.Bytes()))
}
