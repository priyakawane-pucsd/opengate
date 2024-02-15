package dto

import (
	"opengate/models/dao"
	"time"
)

type CreateAuthConfigServiceRequest struct {
	Id         string     `json:"_id,omitempty"`
	Type       string     `json:"type,omitempty"`
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

func (c AuthConfig) toDaoAuthConfig() dao.CreateAuthConfig {
	return dao.CreateAuthConfig{
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

func (r *CreateAuthConfigServiceRequest) ToMongoObject() *dao.AuthConfig {
	return &dao.AuthConfig{
		Id:         r.Id,
		Type:       r.Type,
		AuthConfig: r.AuthConfig.toDaoAuthConfig(),
		CreatedOn:  time.Now().UnixMilli(),
		UpdatedOn:  time.Now().UnixMilli(),
	}
}

type AuthorizationResponse map[string]any
