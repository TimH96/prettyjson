/*
prettyjson.go
author: TimH96
*/

// shell tool to pretty print .json
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	info, err := os.Stdin.Stat()
	// throw piped error
	if err != nil {
		panic(err)
	}
	// panic error when no piped stdin
	if (info.Mode() & os.ModeNamedPipe) == 0 {
		panic("No piped input for prettyjson, use prettyjson --help for usage infomration")
	}
	// parse json input
	reader := bufio.NewReader(os.Stdin)
	//
	fmt.Println(reader.ReadLine())
	fmt.Println(info)
	fmt.Println(err)
}
