package form_test

import (
	"testing"

	"meh/usecase/form"
)

func TestCreateUser_Validate(t *testing.T) {
	type fields struct {
		ScreenName string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "OK",
			fields: fields{
				ScreenName: "gami",
			},
			wantErr: false,
		},
		{
			name: "Empty screen name",
			fields: fields{
				ScreenName: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &form.CreateUser{
				ScreenName: tt.fields.ScreenName,
			}
			if err := f.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("CreateUser.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
