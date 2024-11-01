package dir

import (
	"fmt"
	"io/fs"
	"os"
)

func ListDir(currentDir string) []fs.DirEntry {
	dirFs := os.DirFS(currentDir)

	list, err := fs.ReadDir(dirFs, ".")
	if err != nil {
		fmt.Println("error in reading and listing directory", err)
	}
	return list
}

func MoveToParentDir(currentDir string) []fs.DirEntry {
	parentDir := currentDir + "/.."
	return ListDir(parentDir)
}
