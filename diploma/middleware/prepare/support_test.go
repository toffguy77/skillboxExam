package prepare

import (
	"github.com/toffguy77/statusPage/internal/config"
	"github.com/toffguy77/statusPage/internal/models"
	"reflect"
	"testing"
)

func Test_calcLoad(t *testing.T) {
	config.LoadConfig("../../internal/config/config.json")
	type args struct {
		tickets int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "no tickets",
			args: args{tickets: 0},
			want: 1,
		},
		{
			name: "level 1",
			args: args{tickets: 1},
			want: 1,
		},
		{
			name: "level 2",
			args: args{tickets: 9},
			want: 2,
		},

		{
			name: "level 3",
			args: args{tickets: 16},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcLoad(tt.args.tickets); got != tt.want {
				t.Errorf("calcLoad() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcLoadByTopic(t *testing.T) {
	type args struct {
		supportRes []models.SupportData
	}
	tests := []struct {
		name  string
		args  args
		want  map[string]int
		want1 int
	}{
		{
			name: "default",
			args: args{
				supportRes: []models.SupportData{
					{
						Topic:         "SMS",
						ActiveTickets: 5,
					},
					{
						Topic:         "MMS",
						ActiveTickets: 2,
					},
					{
						Topic:         "Marketing",
						ActiveTickets: 6,
					},
					{
						Topic:         "Other",
						ActiveTickets: 3,
					},
				},
			},
			want: map[string]int{
				"SMS":       5,
				"MMS":       2,
				"Marketing": 6,
				"Other":     3,
			},
			want1: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := calcLoadByTopic(tt.args.supportRes)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcLoadByTopic() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("calcLoadByTopic() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
