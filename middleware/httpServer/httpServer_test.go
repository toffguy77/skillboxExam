package httpServer

import (
	"testing"
)

func Test_isCorrectLoc(t *testing.T) {
	type args struct {
		location string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "default",
			args: args{location: "127.0.0.1:8080"},
			want: true,
		},
		{
			name: "no port",
			args: args{location: "127.0.0.1"},
			want: false,
		},
		{
			name: "no ip",
			args: args{location: ":8080"},
			want: true,
		},
		{
			name: "invalid host:port",
			args: args{location: "asdqwe"},
			want: false,
		},
		{
			name: "ipv6",
			args: args{location: "[::1]:8080"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCorrectLoc(tt.args.location); got != tt.want {
				t.Errorf("isCorrectLoc() = %v, want %v", got, tt.want)
			}
		})
	}
}
