package dto

import (
	"opengate/models/dao"
	"time"
)

type CreateServiceConfigRequest struct {
	Id       string `json:"_id,omitempty"`
	Endpoint string `json:"endpoint"`
	Regex    string `json:"regex"`
}

type GetServiceConfigRequestById struct {
	Id string `json:"_id,omitempty"`
}

func (r *CreateServiceConfigRequest) ToMongoObject() *dao.ServiceConfig {
	return &dao.ServiceConfig{
		Id:        r.Id,
		Endpoint:  r.Endpoint,
		Regex:     r.Regex,
		CreatedOn: time.Now().UnixMilli(),
		UpdatedOn: time.Now().UnixMilli(),
	}
}

type ConfigByIdResponse struct {
	Config     ServiceConfig `json:"config"`
	StatusCode int           `json:"statusCode"`
}

type CreateServiceConfigResponse struct {
	Id         string `json:"_id,omitempty"`
	StatusCode int    `json:"statusCode"`
}

type ListConfigResponse struct {
	Configs    []ServiceConfig `json:"configs"`
	StatusCode int             `json:"statusCode"`
}

type DeleteConfigResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

type ServiceConfig struct {
	Id        string `json:"_id,omitempty"`
	Endpoint  string `json:"endpoint"`
	Regex     string `json:"regex"`
	CreatedOn string `json:"createdOn"`
	UpdatedOn string `json:"updatedOn"`
}
