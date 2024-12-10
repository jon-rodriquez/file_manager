package dir

import (
  "file_manager/cmd/types"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

func MoveToParentDir(currentDir string) (string, []types.Item) {
	splitDir := strings.Split(currentDir, "/")
	if len(splitDir) > 1 {
		splitDir = splitDir[:len(splitDir)-1]
	}
	parentDir := strings.Join(splitDir, "/")
	return parentDir, GetDirItems(parentDir)
}

func GetDirItems(dirPath string) []types.Item {
	items := []types.Item{}
	childDirs := ListDir(dirPath)
	for _, item := range childDirs {
		items = append(items, types.Item{
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
