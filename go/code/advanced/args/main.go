package main

import (
	"fmt"
	"os"
)

// ArgsDemo os.Args 学习demo
func ArgsDemo() {
	fmt.Println("os.Args:", os.Args)
	if len(os.Args) > 0 {
		for i, arg := range os.Args {
			fmt.Printf("args[%d]=%v\n", i, arg)
		}
	}
}

func main() {
	ArgsDemo()
}
