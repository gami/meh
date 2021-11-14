package user

type BlockedUser User

func (u *BlockedUser) Hello() string {
	return "***"
}

type Greetable interface {
	Hello() string
}
