package usecase_test

import (
	"context"
	"errors"
	"sync"
	"testing"

	"meh/core"
	"meh/core/meh"
	"meh/core/user"
	"meh/usecase"
	"meh/usecase/form"

	"github.com/google/go-cmp/cmp"
)

func TestMeh_Meh(t *testing.T) {
	t.Parallel()

	type fields struct {
		tx     core.Tx
		meh    usecase.MehService
		follow usecase.FollowService
	}
	type args struct {
		ctx   context.Context
		input form.CreateMeh
	}

	addToTimelineCalls := make(map[string]struct{})
	wg := sync.WaitGroup{}

	tests := []struct {
		name                   string
		fields                 fields
		args                   args
		want                   meh.ID
		wantErr                bool
		wantAddToTimelineCaled bool
	}{
		{
			name: "OK",
			fields: fields{
				tx: &MockTX{},
				meh: &MockMehService{
					CreateFunc: func(ctx context.Context, meh *meh.Meh) (meh.ID, error) {
						if meh.UserID != 1 {
							return 0, errors.New("MockMeh.Create() invalid userid arg")
						}
						if meh.Text != "test" {
							return 0, errors.New("MockMeh.Create() invalid text arg")
						}

						return 100, nil
					},
					AddToTimelineFunc: func(ctx context.Context, id meh.ID, followeeIDs []user.ID) error {
						if id != 100 {
							return errors.New("MockMeh.AddToTimeline() invalid id arg")
						}

						if len(followeeIDs) != 2 || followeeIDs[0] != 2 || followeeIDs[1] != 1 {
							return errors.New("MockMeh.AddToTimeline() invalid followeeIDs arg")
						}

						addToTimelineCalls["OK"] = struct{}{}
						wg.Done()

						return nil
					},
				},
				follow: &MockFollowService{
					ListFollowersFunc: func(ctx context.Context, userID user.ID) ([]user.ID, error) {
						if userID != 1 {
							return nil, errors.New("MockFollow.ListFollowers() invalid userID arg")
						}

						return []user.ID{
							2,
						}, nil
					},
				},
			},
			args: args{
				ctx: context.Background(),
				input: form.CreateMeh{
					UserID: 1,
					Text:   "test",
				},
			},
			want:                   100,
			wantErr:                false,
			wantAddToTimelineCaled: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := usecase.NewMeh(
				tt.fields.tx,
				tt.fields.meh,
				tt.fields.follow,
			)
			wg.Add(1)
			got, err := u.Meh(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Meh.Meh() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			wg.Wait()

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Meh.Meh() is not match (-got +want):\n%s", diff)
			}

			if _, ok := addToTimelineCalls[tt.name]; ok != tt.wantAddToTimelineCaled {
				t.Errorf("Meh.Meh()'s addToTimelineCalls call result got = %v, wantAddToTimelineCaled %v", ok, tt.wantAddToTimelineCaled)
				return
			}
		})
	}
}
