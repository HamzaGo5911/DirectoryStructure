package main

import (
	"DirectoryStructure/folders"
	"fmt"
)

func main() {
	var numFolders int

	fmt.Print("Enter the number of sub-folders to create ")
	fmt.Scan(&numFolders)

	err := folders.CreateFolders(numFolders)
	if err != nil {
		fmt.Println(err)
		return
	}
}
