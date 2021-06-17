/*
prettyjson.go
author: TimH96
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	info, err := os.Stdin.Stat()
	if err != nil {

	}
	fmt.Println(os.Args)
	fmt.Println(info)
	fmt.Println(err)
}
