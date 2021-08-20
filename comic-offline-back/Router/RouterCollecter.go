package Router

import (
	"comic-offline-back/Security"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	// Resolve the same-origin policy problem
	r.Use(Security.Cors())

	// Apis
	v1 := r.Group("/api/v1")
	{
		GetComicRouter(v1)
		GetStaticFileRouter(v1)
	}

	return r
}
