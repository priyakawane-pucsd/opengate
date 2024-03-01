package config

import (
	"opengate/models/dao"
	"opengate/models/dto"
	"strconv"
)

func convertToListServiceConfigResponse(configs []*dao.Config) []dto.Config {
	var convertedConfigs []dto.Config
	for _, config := range configs {
		convertedConfig := convertToServiceConfigResponse(config)
		convertedConfigs = append(convertedConfigs, *convertedConfig)
	}
	return convertedConfigs
}

func convertToServiceConfigResponse(config *dao.Config) *dto.Config {
	//srvCof := dto.CreateServiceConfigRequest(*config.ServiceConfig)
	return &dto.Config{
		// Map fields from dao.ServiceConfig to dto.ServiceConfigResponse
		Id:            config.Id,
		Type:          config.Type,
		ServiceConfig: &dto.CreateServiceConfigRequest{},
		CreatedOn:     strconv.FormatInt(config.CreatedOn, 10),
		UpdatedOn:     strconv.FormatInt(config.UpdatedOn, 10),
	}
}
