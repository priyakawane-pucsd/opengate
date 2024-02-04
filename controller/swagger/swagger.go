package swagger

import (
	"context"

	_ "opengate/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type SwaggerController struct {
}

func NewSwaggerController(ctx context.Context) *SwaggerController {
	return &SwaggerController{}
}

func (pc *SwaggerController) Register(router gin.IRouter) {
	swaggerRouter := router.Group("/opengate/swagger")
	swaggerRouter.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
