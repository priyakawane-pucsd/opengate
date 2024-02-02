package ping

import "context"

type Service struct {
	repo Repository
}

type Repository interface {
	PingDB(ctx context.Context) error
}

func NewService(ctx context.Context, repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Ping(ctx context.Context) error {
	return s.repo.PingDB(ctx)
}
