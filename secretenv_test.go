package secretenv_test

import (
	"testing"

	"github.com/ndthuan/secretenv"
)

func TestGet(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty key",
			args: args{key: ""},
			want: "",
		},
		{
			name: "secret file found, return its content",
			args: args{key: "MY_SECRET"},
			want: "dxgkscDpDOz0qRadSU/OMrk0nV3DwwzFJmOVA+ei",
		},
		{
			name: "secret file not found, fallback to its corresponding ENV",
			args: args{key: "MY_ENV_VAR"},
			want: "8IZoM8dq2FDkP+yhhGGzFhn+ZaxbHvQDxQNjXoxr",
		},
		{
			name: "neither secret file nor ENV var found, return empty string",
			args: args{key: "NOT_FOUND"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := secretenv.Get(tt.args.key); got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
