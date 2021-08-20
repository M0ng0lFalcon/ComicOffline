package FileController

import (
	"comic-offline-back/Res"
	"comic-offline-back/Service/FileService"
	"github.com/gin-gonic/gin"
)

func GetComicNames(ctx *gin.Context) {
	comicNames := FileService.GetComicNames()
	Res.Success(ctx, gin.H{"comicNames": comicNames}, "get comic names")
}

