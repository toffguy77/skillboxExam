package common

import (
	"testing"
)

func TestIsTrustedProvider(t *testing.T) {
	type args struct {
		provider string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "default",
			args: args{provider: "Topolo"},
			want: true,
		},
		{
			name: "camel case",
			args: args{provider: "kIlDy"},
			want: true,
		},
		{
			name: "invalid",
			args: args{provider: "Topol0"},
			want: false,
		},
		{
			name: "missed",
			args: args{provider: ""},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsTrustedProvider(tt.args.provider); got != tt.want {
				t.Errorf("IsTrustedProvider() = %v, want %v", got, tt.want)
			}
		})
	}
}
