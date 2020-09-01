package main

import (
	"flag"
	"fmt"
	"gopherdojo-studyroom/kadai1/captainblue/imgconv"
)

var (
	beforeExt, afterExt string
)

func init() {
	flag.StringVar(&beforeExt, "b", ".jpg", "set image extension before converted")
	flag.StringVar(&afterExt, "a", ".png", "set image extension after converted")
}

func main() {
	flag.Parse()
	args := flag.Args()
	names, _ := imgconv.GetImgFiles(args[0])
	for _, name := range names {
		fmt.Println(name)
	}
}
