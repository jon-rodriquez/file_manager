package favorites

import (
	"file_manager/cmd/types"
)

func GetFavorties() []types.Item {
	favs := []types.Item{
		{
			Name:  "Documents",
			Path:  "/Users/jonathanrodriquez/Documents",
			IsDir: true,
		},
		{
			Name:  "Downloads",
			Path:  "/Users/jonathanrodriquez/Downloads",
			IsDir: true,
		},
	}
	return favs
}
