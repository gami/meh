package user_test

import (
	"testing"

	"meh/core/user"
)

func TestBlockedUser_Hello(t *testing.T) {
	type obj struct {
		user user.Greetable
	}
	u := &user.User{
		ID:         1,
		ScreenName: "gami",
	}
	tests := []struct {
		name string
		obj  obj
		want string
	}{
		{
			name: "Normal",
			obj: obj{
				user: &user.User{
					ID:         1,
					ScreenName: "gami",
				},
			},
			want: "hello",
		},
		{
			name: "Normal",
			obj: obj{
				user: (*user.BlockedUser)(u),
			},
			want: "***",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.obj.user
			if got := s.Hello(); got != tt.want {
				t.Errorf("User.Hello() = %v, want %v", got, tt.want)
			}
		})
	}
}
