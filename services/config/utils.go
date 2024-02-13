package config

import (
	"opengate/models/dao"
	"opengate/models/dto"
	"strconv"
)

func convertToListServiceConfigResponse(configs []*dao.ServiceConfig) []dto.ServiceConfig {
	var convertedConfigs []dto.ServiceConfig
	for _, config := range configs {
		convertedConfig := convertToServiceConfigResponse(config)
		convertedConfigs = append(convertedConfigs, *convertedConfig)
	}
	return convertedConfigs
}

func convertToServiceConfigResponse(config *dao.ServiceConfig) *dto.ServiceConfig {
	return &dto.ServiceConfig{
		// Map fields from dao.ServiceConfig to dto.ServiceConfigResponse
		Id:            config.Id,
		Type:          config.Type,
		ServiceConfig: dto.CreateServiceConfigRequest(config.ServiceConfig),
		CreatedOn:     strconv.FormatInt(config.CreatedOn, 10),
		UpdatedOn:     strconv.FormatInt(config.UpdatedOn, 10),
	}
}

func convertToServiceConfigByIdResponse(config *dao.ServiceConfig) *dto.ServiceConfig {
	return &dto.ServiceConfig{
		// Map fields from dao.ServiceConfig to dto.ServiceConfigResponse
		Id:            config.Id,
		Type:          config.Type,
		ServiceConfig: dto.CreateServiceConfigRequest(config.ServiceConfig),
		CreatedOn:     strconv.FormatInt(config.CreatedOn, 10),
		UpdatedOn:     strconv.FormatInt(config.UpdatedOn, 10),
	}
}
