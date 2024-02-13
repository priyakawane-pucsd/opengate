package dto

import (
	"opengate/models/dao"
	"time"
)

type CreateAuthConfigServiceRequest struct {
	Id         string            `json:"_id,omitempty"`
	Type       string            `json:"type,omitempty"`
	AuthConfig AuthConfigRequest `json:"authConfig1,omitempty"`
}

type AuthConfigRequest struct {
	Endpoint string   `json:"endpoint"`
	Headers  []string `json:"headers"` //new
}

type CreateAuthConfigResponse struct {
	Id         string `json:"_id,omitempty"`
	StatusCode int    `json:"statusCode"`
}

func (r *CreateAuthConfigServiceRequest) ToMongoObject() *dao.AuthConfigService {
	return &dao.AuthConfigService{
		Id:         r.Id,
		Type:       r.Type,
		AuthConfig: dao.CreateAuthConfigRequest1(r.AuthConfig),
		CreatedOn:  time.Now().UnixMilli(),
		UpdatedOn:  time.Now().UnixMilli(),
	}
}
