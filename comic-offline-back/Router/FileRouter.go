package Router

import (
	"comic-offline-back/Controller/FileController"
	"github.com/gin-gonic/gin"
)

func GetFileRouter(v1 *gin.RouterGroup) *gin.RouterGroup {
	FileRouter := v1.Group("/file")
	{
		FileRouter.GET("/comicNames", FileController.GetComicNames)
	}

	return FileRouter
}