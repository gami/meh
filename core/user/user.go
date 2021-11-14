package user

type ID uint64
type IDs []ID
type idKey struct{}

var IDContextKey idKey

type User struct {
	ID         ID
	ScreenName string
}

func (u *User) Hello() string {
	return "hello"
}

func (is IDs) ToUint64s() []uint64 {
	var res []uint64

	for _, i := range is {
		res = append(res, uint64(i))
	}

	return res
}
