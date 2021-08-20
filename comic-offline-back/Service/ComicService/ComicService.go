package ComicService

import (
	"comic-offline-back/Utils"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
)

func GetComicNames() []string {
	ROOT := viper.GetString("comic.root")
	dirList, err := ioutil.ReadDir(ROOT)
	if err != nil {
		msg := "[!] Can't Find file path"
		log.Println(msg)
	}

	// return comic names
	comicNames := Utils.GetDirNames(dirList)

	return comicNames
}

func GetComicChapters(ComicName string) []string {
	// make path
	ROOT := viper.GetString("comic.root")
	ComicPath := fmt.Sprintf("%s/%s", ROOT, ComicName)

	dirList, err := ioutil.ReadDir(ComicPath)
	if err != nil {
		msg := "[!] Can't Find file path"
		log.Println(msg)
	}

	Chapters := Utils.GetDirNames(dirList)

	return Chapters
}

func GetChapterPages(ComicName string, Chapter string) ([]string, int) {
	ROOT := viper.GetString("comic.root")
	ChapterPath := fmt.Sprintf("%s/%s/%s", ROOT, ComicName, Chapter)

	dirList, err := ioutil.ReadDir(ChapterPath)
	if err != nil {
		msg := "[!] Can't find this chapter"
		log.Println(msg)
	}

	UrlRoot := "comicRoot"

	pageNames := Utils.GetFileNames(dirList)
	pageUrls := make([]string, 0)
	for _, v := range pageNames {
		// static file path, comic name, chapter, page name
		url := fmt.Sprintf("/%s/%s/%s/%s", UrlRoot, ComicName, Chapter, v)
		pageUrls = append(pageUrls, url)
	}

	return pageUrls, len(pageUrls)
}
