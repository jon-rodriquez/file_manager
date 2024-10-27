package dir

import (
	"fmt"
	"io/fs"
	"os"
)

func ListDir(dir string) {
	dirFs := os.DirFS(dir)
	list, err := fs.ReadDir(dirFs, ".")
	if err == nil {
		for _, item := range list {
			fmt.Println(item)
		}
	}
}
