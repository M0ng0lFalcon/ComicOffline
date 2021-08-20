package FileService

import (
	"io/ioutil"
)

func GetComicNames() []string {
	comicNames := make([]string, 0)

	ROOT := "./Resources/comic_books"
	dirList, err := ioutil.ReadDir(ROOT)
	if err != nil {
		panic(err)
	}
	for _, v := range dirList {
		if v.IsDir() {
			comicNames = append(comicNames, v.Name())
		}
	}

	return comicNames
}