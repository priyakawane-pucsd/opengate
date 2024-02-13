package authconfig

import (
	"context"
	"opengate/models/dto"
	"opengate/utils"

	"github.com/bappaapp/goutils/logger"
	"github.com/gin-gonic/gin"
)

type Config struct {
	DB bool
}

type AuthConfigController struct {
	cfg     *Config
	service Service
}

type Service interface {
	CreateUpdateAuthConfig(ctx context.Context, req *dto.CreateAuthConfigServiceRequest) (*dto.CreateAuthConfigResponse, error)
}

func NewAuthConfigController(ctx context.Context, cfg *Config, s Service) *AuthConfigController {
	return &AuthConfigController{cfg: cfg, service: s}
}

func (ac *AuthConfigController) Register(router gin.IRouter) {
	configRouter := router.Group("/opengate/authConfig")
	configRouter.PUT("/", ac.CreateUpdateAuthConfig)
}

// CreateUpdateConfig handles the creation or update of a service configuration.
// @Summary Create or update service configuration.
// @Description Create or update service configuration based on the provided request.
// @Tags AuthConfig
// @Accept json
// @Produce json
// @Param request body dto.CreateAuthConfigServiceRequest true "Request body containing configuration details"
// @Success 200 {object} dto.CreateAuthConfigResponse "Successful response"
// @Failure 400 {object} utils.CustomError "Invalid request body"
// @Failure 500 {object} utils.CustomError "Internal server error"
// @Router /opengate/authConfig [put]
func (ac *AuthConfigController) CreateUpdateAuthConfig(ctx *gin.Context) {
	var req dto.CreateAuthConfigServiceRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		logger.Error(ctx, "Invalid Body: %s", err.Error())
		utils.WriteError(ctx, utils.NewBadRequestError("Invalid body"))
		return
	}
	res, err := ac.service.CreateUpdateAuthConfig(ctx, &req)
	if err != nil {
		utils.WriteError(ctx, err)
		return
	}
	utils.WriteResponse(ctx, res)
}
