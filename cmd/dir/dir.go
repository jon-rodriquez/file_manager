package dir

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
)

func ListDir(currentDir string) []fs.DirEntry {
	dirFs := os.DirFS(currentDir)

	list, err := fs.ReadDir(dirFs, ".")
	if err != nil {
		fmt.Println("error in reading and listing directory", err)
	}
	return list
}

func MoveToParentDir(currentDir string) (string, []fs.DirEntry) {
  splitDir := strings.Split(currentDir, "/")
  if len(splitDir) > 1 {
    splitDir = splitDir[:len(splitDir)-1]
  }
  parentDir := strings.Join(splitDir, "/")
	return parentDir, ListDir(parentDir)
}
