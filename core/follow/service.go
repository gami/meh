package follow

import (
	"context"
	"log"

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

func (s *Service) Follow(ctx context.Context, from, to user.ID) error {
	if from == to {
		return errors.New("cannot follow same id")
	}

	found, err := s.repo.Exists(ctx, from, to)
	if err != nil {
		return errors.Wrapf(err, "failed to check follow from=%v to=%v", from, to)
	}

	if found {
		log.Printf("follow is ignored because alreday followed from=%v to=%v", from, to)
		return nil
	}

	return s.repo.Create(ctx, from, to)
}

func (s *Service) Remove(ctx context.Context, from, to user.ID) error {
	found, err := s.repo.Exists(ctx, from, to)
	if err != nil {
		return errors.Wrapf(err, "failed to check follow from=%v to=%v", from, to)
	}

	if !found {
		log.Printf("remove is ignored because not followed from=%v to=%v", from, to)
		return nil
	}

	return s.repo.Delete(ctx, from, to)
}

func (s *Service) ListFollowers(ctx context.Context, userID user.ID) ([]user.ID, error) {
	return s.repo.ListFollowers(ctx, userID)
}
