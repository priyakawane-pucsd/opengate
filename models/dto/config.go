package dto

import (
	"opengate/models/dao"
	"time"
)

type CreateConfigServiceRequest struct {
	Id            string                     `json:"_id,omitempty"`
	Type          string                     `json:"type,omitempty"`
	ServiceConfig CreateServiceConfigRequest `json:"serviceConfig,omitempty"`
}

type CreateServiceConfigRequest struct {
	Name     string        `json:"name"`
	Endpoint string        `json:"endpoint"`
	Regex    string        `json:"regex"`
	Apis     []ServiceApis `json:"apis"`
}

func (sc *CreateServiceConfigRequest) ToDaoObject() *dao.ServiceConfig {
	var apis []dao.ServiceApis
	for _, v := range sc.Apis {
		apis = append(apis, dao.ServiceApis{
			Authorization: v.Authorization,
			Regex:         v.Regex,
			Roles:         v.Roles,
		})
	}
	return &dao.ServiceConfig{
		Name:     sc.Name,
		Endpoint: sc.Endpoint,
		Regex:    sc.Regex,
		Apis:     apis,
	}
}

type ServiceApis struct {
	Regex         string   `json:"regex"`
	Authorization bool     `json:"authorization"`
	Roles         []string `json:"roles"`
}

type GetServiceConfigRequestById struct {
	Id string `json:"_id,omitempty"`
}

func (r *CreateConfigServiceRequest) ToMongoObject() *dao.Config {
	return &dao.Config{
		Id:            r.Id,
		Type:          r.Type,
		ServiceConfig: r.ServiceConfig.ToDaoObject(),
		CreatedOn:     time.Now().UnixMilli(),
		UpdatedOn:     time.Now().UnixMilli(),
	}
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

type ConfigByIdResponse struct {
	Config     Config `json:"config"`
	StatusCode int    `json:"statusCode"`
}
