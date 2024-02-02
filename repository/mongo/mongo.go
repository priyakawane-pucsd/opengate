package mongo

import "context"

type Repository struct {
}

func NewRepository(ctx context.Context, config *Config) *Repository {
	return &Repository{}
}

func (r *Repository) PingDB(ctx context.Context) error {
	//todo
	return nil
}
