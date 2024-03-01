package config

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
	CreateUpdateConfig(ctx context.Context, cfg *dao.Config) error
	GetAllConfigs(ctx context.Context) ([]*dao.Config, error)
	GetConfigById(ctx context.Context, id string) (*dao.Config, error)
	DeleteConfigById(ctx context.Context, id string) error
}

func NewService(ctx context.Context, repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUpdateConfig(ctx context.Context, req *dto.CreateConfigServiceRequest) (*dto.CreateConfigServiceResponse, error) {
	cfg := req.ToMongoObject()
	err := s.repo.CreateUpdateConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}
	return &dto.CreateConfigServiceResponse{StatusCode: http.StatusCreated, Id: cfg.Id}, nil
}

func (s *Service) GetAllConfigs(ctx context.Context) (*dto.ListConfigResponse, error) {
	configs, err := s.repo.GetAllConfigs(ctx)
	if err != nil {
		return nil, err
	}

	response := &dto.ListConfigResponse{
		Configs:    convertToListServiceConfigResponse(configs), // Assuming configs is a field in the ResultConfigResponse
		StatusCode: http.StatusOK,
	}
	return response, nil
}

func (s *Service) GetConfigById(ctx context.Context, id string) (*dto.ConfigByIdResponse, error) {
	config, err := s.repo.GetConfigById(ctx, id)
	if err != nil {
		return nil, err
	}

	response := &dto.ConfigByIdResponse{
		Config:     *convertToServiceConfigResponse(config),
		StatusCode: http.StatusOK,
	}
	return response, nil
}

func (s *Service) DeleteConfigById(ctx context.Context, id string) (*dto.DeleteConfigResponse, error) {
	err := s.repo.DeleteConfigById(ctx, id)
	if err != nil {
		return nil, err
	}

	// If the deletion is successful, create a response
	response := &dto.DeleteConfigResponse{
		Message:    "Configuration deleted successfully",
		StatusCode: http.StatusOK,
	}
	return response, nil
}

func (s *Service) CreateUpdateAuthConfig(ctx context.Context, req *dto.CreateAuthConfigServiceRequest) (*dto.CreateAuthConfigResponse, error) {
	cfg := req.ToMongoObject()
	err := s.repo.CreateUpdateConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}
	return &dto.CreateAuthConfigResponse{StatusCode: http.StatusCreated, Id: cfg.Id}, nil
}
