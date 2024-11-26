package dir

func GetFavorties() []Item {
	favs := []Item{
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
