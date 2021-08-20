package Router

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func GetStaticFileRouter(v1 *gin.RouterGroup) *gin.RouterGroup {
	// static file system
	ComicRoot := viper.GetString("comic.root")
	v1.Static("comicRoot", ComicRoot)
	return v1
}