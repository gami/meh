package meh

import (
	"context"

	"meh/core"
	"meh/core/user"

	"github.com/pkg/errors"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(ctx context.Context, meh *Meh) (ID, error) {
	return s.repo.Create(ctx, meh)
}

func (s *Service) AddToTimeline(ctx context.Context, id ID, followeeIDs []user.ID) error {
	// TODO 20-100件くらいずつbulkで追加した方が良さそう
	for _, f := range followeeIDs {
		err := s.repo.AddToTimeline(ctx, id, f)
		if err != nil {
			return errors.Wrap(err, "failed to add to timeline")
		}
	}

	return nil
}

func (s *Service) ListMehsInTimeline(ctx context.Context, userID user.ID, pagination core.Pagination) (Mehs, core.Pagination, error) {
	return s.repo.ListMehsInTimeline(ctx, userID, pagination)
}
