package follow

import (
	"context"

	"meh/core/user"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Follow(ctx context.Context, from, to user.ID) error {
	return s.repo.Create(ctx, from, to)
}

func (s *Service) Remove(ctx context.Context, from, to user.ID) error {
	return s.repo.Delete(ctx, from, to)
}

func (s *Service) Exists(ctx context.Context, from, to user.ID) (bool, error) {
	return s.repo.Exists(ctx, from, to)
}

func (s *Service) ListFollowers(ctx context.Context, userID user.ID) ([]user.ID, error) {
	return s.repo.ListFollowers(ctx, userID)
}
