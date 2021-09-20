package di

import (
	"meh/core/follow"
	"meh/core/meh"
	"meh/core/user"
	"meh/repository"
)

func InjectUserRepository() user.Repository {
	return repository.NewUser(
		InjectEntClient(),
	)
}

func InjectFollowRepository() follow.Repository {
	return repository.NewFollow(
		InjectEntClient(),
	)
}

func InjectMehRepository() meh.Repository {
	return repository.NewMeh(
		InjectEntClient(),
	)
}
