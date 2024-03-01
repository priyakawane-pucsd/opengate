package dto

import (
	"opengate/models/dao"
	"time"
)

// 13-02-2024 Newly added
type CreateConfigServiceRequest struct {
	Id            string                     `json:"_id,omitempty"`
	Type          string                     `json:"type,omitempty"`
	ServiceConfig CreateServiceConfigRequest `json:"serviceConfig,omitempty"`
}

type CreateServiceConfigRequest struct {
	Endpoint      string   `json:"endpoint"`
	Regex         string   `json:"regex"`
	Authorization bool     `json:"authorization"` //new
	Roles         []string `json:"roles"`         //new
}

type GetServiceConfigRequestById struct {
	Id string `json:"_id,omitempty"`
}

// new added
func (r *CreateConfigServiceRequest) ToMongoObject() *dao.Config {
	serConf := dao.ServiceConfig(r.ServiceConfig)
	return &dao.Config{
		Id:            r.Id,
		Type:          r.Type,
		ServiceConfig: &serConf,
		CreatedOn:     time.Now().UnixMilli(),
		UpdatedOn:     time.Now().UnixMilli(),
	}
}

type ConfigByIdResponse struct {
	Config     Config `json:"config"`
	StatusCode int    `json:"statusCode"`
}

type CreateConfigServiceResponse struct {
	Id         string `json:"_id,omitempty"`
	StatusCode int    `json:"statusCode"`
}

type ListConfigResponse struct {
	Configs    []Config `json:"configs"`
	StatusCode int      `json:"statusCode"`
}

type DeleteConfigResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

type Config struct {
	Id            string                      `json:"_id,omitempty"`
	Type          string                      `json:"type,omitempty"`
	ServiceConfig *CreateServiceConfigRequest `json:"serviceConfig,omitempty"`
	AuthConfig    *AuthConfig                 `json:"authConfig,omitempty"`
	CreatedOn     string                      `json:"createdOn"`
	UpdatedOn     string                      `json:"updatedOn"`
}
