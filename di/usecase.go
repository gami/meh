package di

import (
	"meh/controller"
	"meh/usecase"
)

func InjectUserUsecase() controller.UserUsecase {
	return usecase.NewUser(
		InjectTx(),
		InjectUserService(),
	)
}

func InjectFollowUsecase() controller.FollowUsecase {
	return usecase.NewFollow(
		InjectTx(),
		InjectFollowService(),
	)
}

func InjectMehUsecase() controller.MehUsecase {
	return usecase.NewMeh(
		InjectTx(),
		InjectMehService(),
		InjectFollowService(),
	)
}
