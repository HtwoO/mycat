package main

import "os"

func main() {
	switch len(os.Args) {
	case 1: // no argument provided
		usage()
	}
}

func usage() {
	println(os.Args[0], "Usage 使用说明")
}
