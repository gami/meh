package usecase_test

import (
	"context"

	"meh/core"
	"meh/core/meh"
	"meh/core/user"
)

type MockTX struct{}

func (m *MockTX) Transact(ctx context.Context, f func(context.Context) (interface{}, error)) (interface{}, error) {
	return f(ctx)
}

type MockMehService struct {
	CreateFunc             func(ctx context.Context, meh *meh.Meh) (meh.ID, error)
	AddToTimelineFunc      func(ctx context.Context, id meh.ID, followeeIDs []user.ID) error
	ListMehsInTimelineFunc func(ctx context.Context, userID user.ID, pagination core.Pagination) (meh.Mehs, core.Pagination, error)
}

func (m *MockMehService) Create(ctx context.Context, mh *meh.Meh) (meh.ID, error) {
	return m.CreateFunc(ctx, mh)
}
func (m *MockMehService) AddToTimeline(ctx context.Context, id meh.ID, followeeIDs []user.ID) error {
	return m.AddToTimelineFunc(ctx, id, followeeIDs)
}
func (m *MockMehService) ListMehsInTimeline(ctx context.Context, userID user.ID, pagination core.Pagination) (meh.Mehs, core.Pagination, error) {
	return m.ListMehsInTimelineFunc(ctx, userID, pagination)
}

type MockFollowService struct {
	FollowFunc        func(ctx context.Context, from, to user.ID) error
	RemoveFunc        func(ctx context.Context, from, to user.ID) error
	ListFollowersFunc func(ctx context.Context, userID user.ID) ([]user.ID, error)
}

func (m *MockFollowService) Follow(ctx context.Context, from, to user.ID) error {
	return m.FollowFunc(ctx, from, to)
}
func (m *MockFollowService) Remove(ctx context.Context, from, to user.ID) error {
	return m.RemoveFunc(ctx, from, to)
}
func (m *MockFollowService) ListFollowers(ctx context.Context, userID user.ID) ([]user.ID, error) {
	return m.ListFollowersFunc(ctx, userID)
}
