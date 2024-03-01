package dto

import (
	"opengate/constants"
	"opengate/models/dao"
	"time"
)

type CreateAuthConfigServiceRequest struct {
	AuthConfig AuthConfig `json:"authConfig,omitempty"`
}

type AuthConfig struct {
	Endpoint       string          `json:"endpoint"`
	Headers        []string        `json:"headers"`
	RequestMethod  string          `json:"requestMethod"`
	ForwardHeaders []ForwardHeader `json:"forwardHeaders"`
	RolesKey       string          `json:"rolesKey"`
}

func dtoToDaoForwardHeaders(header []ForwardHeader) []dao.ForwardHeader {
	var hds []dao.ForwardHeader
	for _, h := range header {
		hds = append(hds, dao.ForwardHeader{Key: h.Key, Address: h.Address})
	}
	return hds
}

func (c AuthConfig) toDaoAuthConfig() dao.AuthConfig {
	return dao.AuthConfig{
		Endpoint:       c.Endpoint,
		Headers:        c.Headers,
		RequestMethod:  c.RequestMethod,
		ForwardHeaders: dtoToDaoForwardHeaders(c.ForwardHeaders),
		RolesKey:       c.RolesKey,
	}
}

// ForwardHeader represents the structure of the ForwardHeaders array
type ForwardHeader struct {
	Key     string `json:"key"`
	Address string `json:"address"`
}

type CreateAuthConfigResponse struct {
	Id         string `json:"_id,omitempty"`
	StatusCode int    `json:"statusCode"`
}

func (r *CreateAuthConfigServiceRequest) ToMongoObject() *dao.Config {
	authConf := r.AuthConfig.toDaoAuthConfig()
	return &dao.Config{
		Id:         constants.AUTH_CONFIG,
		Type:       constants.AUTH_CONFIG,
		AuthConfig: &authConf,
		CreatedOn:  time.Now().UnixMilli(),
		UpdatedOn:  time.Now().UnixMilli(),
	}
}

type AuthorizationResponse map[string]any
