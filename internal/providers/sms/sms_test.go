package sms

import (
	"github.com/toffguy77/statusPage/internal/models"
	"reflect"
	"testing"
)

func Test_parseString(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want models.SMSData
	}{
		{
			name: "default",
			args: args{line: "RU;1;499;Topolo"},
			want: models.SMSData{
				Country:      "RU",
				Bandwidth:    "1",
				ResponseTime: "499",
				Provider:     "Topolo",
			},
		},
		{
			name: "fewer fields",
			args: args{line: "RU;499;Topolo"},
			want: models.SMSData{},
		},
		{
			name: "more fields",
			args: args{line: "RU;1;499;Topolo;Alive"},
			want: models.SMSData{},
		},
		{
			name: "not trusted provider",
			args: args{line: "RU;1;499;Beeline"},
			want: models.SMSData{},
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
