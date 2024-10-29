package dir

import (
	"fmt"
	"io/fs"
	"os"
)

func ListDir(dir string) []fs.DirEntry {
	dirFs := os.DirFS(dir)
	list, err := fs.ReadDir(dirFs, ".")
	if err != nil {
		fmt.Println("error in reading and listing directory")
	}
	return list
}
