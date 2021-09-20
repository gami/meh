package di

import (
	openapi "meh/api/openapi"
	"meh/controller"
)

func InjectController() openapi.ServerInterface {
	return controller.NewController(
		InjectUserController(),
		InjectFollowController(),
		InjectMehController(),
	)
}

func InjectUserController() *controller.User {
	return controller.NewUser(
		InjectUserUsecase(),
	)
}

func InjectFollowController() *controller.Follow {
	return controller.NewFollow(
		InjectFollowUsecase(),
	)
}

func InjectMehController() *controller.Meh {
	return controller.NewMeh(
		InjectMehUsecase(),
		InjectUserUsecase(),
	)
}
