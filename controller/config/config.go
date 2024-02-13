package config

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

type ConfigController struct {
	cfg     *Config
	service Service
}

type Service interface {
	// CreateUpdateConfig(ctx context.Context, req *dto.CreateServiceConfigRequest) (*dto.CreateServiceConfigResponse, error)
	CreateUpdateConfig(ctx context.Context, req *dto.CreateConfigServiceRequest) (*dto.CreateConfigServiceResponse, error)
	GetAllConfigs(ctx context.Context) (*dto.ListConfigResponse, error)
	GetConfigById(ctx context.Context, id string) (*dto.ConfigByIdResponse, error)
	DeleteConfigById(ctx context.Context, id string) (*dto.DeleteConfigResponse, error)
}

func NewConfigController(ctx context.Context, cfg *Config, s Service) *ConfigController {
	return &ConfigController{cfg: cfg, service: s}
}

func (cc *ConfigController) Register(router gin.IRouter) {
	configRouter := router.Group("/opengate/config")
	configRouter.PUT("/", cc.CreateUpdateConfig)
	configRouter.GET("/", cc.GetAllConfigs)
	configRouter.GET("/:id", cc.GetConfigById)
	configRouter.DELETE("/:id", cc.DeleteConfigById)
}

// CreateUpdateConfig handles the creation or update of a service configuration.
// @Summary Create or update service configuration.
// @Description Create or update service configuration based on the provided request.
// @Tags Config
// @Accept json
// @Produce json
// @Param request body dto.CreateConfigServiceRequest true "Request body containing configuration details"
// @Success 200 {object} dto.CreateConfigServiceResponse "Successful response"
// @Failure 400 {object} utils.CustomError "Invalid request body"
// @Failure 500 {object} utils.CustomError "Internal server error"
// @Router /opengate/config/ [put]
func (cc *ConfigController) CreateUpdateConfig(ctx *gin.Context) {
	var req dto.CreateConfigServiceRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		logger.Error(ctx, "Invalid Body: %s", err.Error())
		utils.WriteError(ctx, utils.NewBadRequestError("Invalid body"))
		return
	}
	res, err := cc.service.CreateUpdateConfig(ctx, &req)
	if err != nil {
		utils.WriteError(ctx, err)
		return
	}
	utils.WriteResponse(ctx, res)
}

// GetAllConfigs retrieves all service configurations.
// @Summary Get all service configurations
// @Description Retrieve a list of all service configurations.
// @Tags Config
// @Accept json
// @Produce json
// @Success 200 {array} dto.ListConfigResponse "Successful response"
// @Failure 400 {object} utils.CustomError "Invalid request"
// @Failure 500 {object} utils.CustomError "Internal server error"
// @Router /opengate/config/ [get]
func (cc *ConfigController) GetAllConfigs(ctx *gin.Context) {
	configs, err := cc.service.GetAllConfigs(ctx)
	if err != nil {
		utils.WriteError(ctx, err)
		return
	}
	utils.WriteResponse(ctx, configs)
}

// GetConfigById retrieves a service configuration by ID.
// @Summary Get a service configuration by ID.
// @Description Retrieve a service configuration based on the provided request containing the ID.
// @Tags Config
// @Accept json
// @Produce json
// @Param id path string true "ID of the service configuration"
// @Success 200 {object} dto.ConfigByIdResponse "Successful response"
// @Failure 400 {object} utils.CustomError "Invalid request body"
// @Failure 404 {object} utils.CustomError "Config not found"
// @Failure 500 {object} utils.CustomError "Internal server error"
// @Router /opengate/config/{id} [get]
func (cc *ConfigController) GetConfigById(ctx *gin.Context) {
	id := ctx.Param("id")
	res, err := cc.service.GetConfigById(ctx, id)
	if err != nil {
		utils.WriteError(ctx, err)
		return
	}
	utils.WriteResponse(ctx, res)
}

// DeleteConfigById deletes a service configuration by ID.
// @Summary Delete a service configuration by ID.
// @Description Delete a service configuration based on the provided ID.
// @Tags Config
// @Accept json
// @Produce json
// @Param id path string true "ID of the service configuration to delete"
// @Success 200 {object} dto.DeleteConfigResponse "Successful response"
// @Failure 400 {object} utils.CustomError "Invalid ID format"
// @Failure 404 {object} utils.CustomError "Config not found"
// @Failure 500 {object} utils.CustomError "Internal server error"
// @Router /opengate/config/{id} [delete]
func (cc *ConfigController) DeleteConfigById(ctx *gin.Context) {
	id := ctx.Param("id")
	res, err := cc.service.DeleteConfigById(ctx, id)
	if err != nil {
		utils.WriteError(ctx, err)
		return
	}
	utils.WriteResponse(ctx, res)
}
