package Router

import (
	"comic-offline-back/Controller/ComicController"
	"github.com/gin-gonic/gin"
)

func GetComicRouter(v1 *gin.RouterGroup) *gin.RouterGroup {
	ComicRouter := v1.Group("/comic")
	{
		ComicRouter.GET("/comicNames", ComicController.GetComicNames)
		ComicRouter.GET("/comicChapters", ComicController.GetComicChapters)
		ComicRouter.GET("/chapterPages", ComicController.GetChapterPages)
	}

	return ComicRouter
}