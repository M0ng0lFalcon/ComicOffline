package ComicController

import (
	"comic-offline-back/Res"
	"comic-offline-back/Service/ComicService"
	"github.com/gin-gonic/gin"
)

func GetComicNames(ctx *gin.Context) {
	comicNames := ComicService.GetComicNames()
	Res.Success(ctx, gin.H{"comicNames": comicNames}, "get comic names")
}

func GetComicChapters(ctx *gin.Context) {
	// get comic name
	ComicName := ctx.Query("ComicName")

	Chapters := ComicService.GetComicChapters(ComicName)

	Res.Success(ctx, gin.H{"Chapters": Chapters}, "get comic chapters")
}

func GetChapterPages(ctx *gin.Context) {
	// get comic name and chapter name
	ComicName := ctx.Query("ComicName")
	Chapter := ctx.Query("Chapter")

	pageUrls, pageCount := ComicService.GetChapterPages(ComicName, Chapter)

	Res.Success(ctx, gin.H{"pageUrls": pageUrls, "pageCount": pageCount}, "get chapter pages")
}
