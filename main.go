package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// fmt.Println("Number of arguments:", len(os.Args))
	// fmt.Println("Arguments:", os.Args)

	if len(os.Args) != 3 {
		fmt.Println("only 2 args are accepted received ", len(os.Args)-1)
	}

	sourceDir := os.Args[1]
	destDir := os.Args[2]

	fmt.Println("Source Input => ", sourceDir)
	fmt.Println("Destination Input => ", destDir)

	switch destDir {
	case "../":
		countInput := len(strings.Split(destDir, "/"))

		tempSrcDir := strings.Split(sourceDir, "/")

		fmt.Println("countInput => ", countInput)
		fmt.Println("len(tempSrcDir) => ", len(tempSrcDir))

		if countInput > len(tempSrcDir) {
			destDir = "/"

		} else if countInput == len(tempSrcDir) {
			destDir = "/home"

		} else {
			destDir = strings.Join(tempSrcDir[:len(tempSrcDir)-countInput], "/")

		}

	case "./":
		destDir = sourceDir
	}

	fmt.Println("Source Output => ", sourceDir)
	fmt.Println("Destination Output => ", destDir)
}