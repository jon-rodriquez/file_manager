package dir

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
)

type Item struct {
	Name  string
	Path  string
	IsDir bool
}

func MoveToParentDir(currentDir string) (string, []Item) {
	splitDir := strings.Split(currentDir, "/")
	if len(splitDir) > 1 {
		splitDir = splitDir[:len(splitDir)-1]
	}
	parentDir := strings.Join(splitDir, "/")
	return parentDir, GetDirItems(parentDir)
}

func GetDirItems(dirPath string) []Item {
	items := []Item{}
	childDirs := ListDir(dirPath)
	for _, item := range childDirs {
		items = append(items, Item{
			Name:  item.Name(),
			Path:  item.Name(),
			IsDir: item.IsDir(),
		})
	}
	return items
}

func ListDir(currentDir string) []fs.DirEntry {
	dirFs := os.DirFS(currentDir)

	list, err := fs.ReadDir(dirFs, ".")
	if err != nil {
		fmt.Println("error in reading and listing directory", err)
	}
	return list
}
