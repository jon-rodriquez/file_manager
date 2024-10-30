package dir

import (
	"fmt"
	"io/fs"
	"os"
)

func ListDir() []fs.DirEntry {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("error in getting current directory", err)
		return nil
	}
  dirFs := os.DirFS(currentDir)

	list, err := fs.ReadDir(dirFs, ".")
	if err != nil {
		fmt.Println("error in reading and listing directory", err)
	}
	return list
}
