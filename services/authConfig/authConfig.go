package authconfig

import (
	"context"
	"net/http"
	"opengate/models/dao"
	"opengate/models/dto"
)

type Service struct {
	repo Repository
}

type Repository interface {
	CreateUpdateAuthConfig(ctx context.Context, cfg *dao.AuthConfig) error
}

func NewService(ctx context.Context, repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUpdateAuthConfig(ctx context.Context, req *dto.CreateAuthConfigServiceRequest) (*dto.CreateAuthConfigResponse, error) {
	cfg := req.ToMongoObject()
	err := s.repo.CreateUpdateAuthConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}
	return &dto.CreateAuthConfigResponse{StatusCode: http.StatusCreated, Id: cfg.Id}, nil
}
