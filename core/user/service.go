package user

import (
	"context"

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

func (s *Service) Create(ctx context.Context, user *User) (ID, error) {
	found, err := s.repo.ExistsByScreenName(ctx, user.ScreenName)
	if err != nil {
		return 0, errors.Wrap(err, "failed to check screen name")
	}

	if found {
		return 0, errors.New("screen name is already used")
	}

	id, err := s.repo.Create(ctx, user)
	if err != nil {
		return 0, errors.Wrap(err, "failed to create user")
	}

	return id, nil
}

func (s *Service) FindByUserIDs(ctx context.Context, userIDs IDs) ([]*User, error) {
	return s.repo.FindByUserIDs(ctx, userIDs)
}
