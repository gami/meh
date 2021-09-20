package di

import (
	"meh/core/follow"
	"meh/core/meh"
	"meh/core/user"
	"meh/usecase"
)

func InjectUserService() usecase.UserService {
	return user.NewService(
		InjectUserRepository(),
	)
}

func InjectFollowService() usecase.FollowService {
	return follow.NewService(
		InjectFollowRepository(),
	)
}

func InjectMehService() usecase.MehService {
	return meh.NewService(
		InjectMehRepository(),
	)
}
