package email

import (
	"github.com/toffguy77/statusPage/internal/models"
	"reflect"
	"testing"
)

func Test_isTrustedEmailProvider(t *testing.T) {
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
			args: args{provider: "gmail"},
			want: true,
		},
		{
			name: "upper case",
			args: args{provider: "MAIL.RU"},
			want: true,
		},
		{
			name: "untrusted",
			args: args{provider: "rambler"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isTrustedEmailProvider(tt.args.provider); got != tt.want {
				t.Errorf("isTrustedEmailProvider() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseString(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want models.EmailData
	}{
		{
			name: "default",
			args: args{line: "RU;Yandex;210"},
			want: models.EmailData{
				Country:      "RU",
				Provider:     "Yandex",
				DeliveryTime: 210,
			},
		},
		{
			name: "empty field",
			args: args{line: "RU;;210"},
			want: models.EmailData{},
		},
		{
			name: "missed field",
			args: args{line: "RU;Yandex"},
			want: models.EmailData{},
		},
		{
			name: "mode fields",
			args: args{line: "RU;Yandex;210;Alive"},
			want: models.EmailData{},
		},
		{
			name: "untrusted provider",
			args: args{line: "RU;Rambler;210"},
			want: models.EmailData{},
		},
		{
			name: "delivery is not an int",
			args: args{line: "RU;Yandex;210.1"},
			want: models.EmailData{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseString(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseString() = %v, want %v", got, tt.want)
			}
		})
	}
}
