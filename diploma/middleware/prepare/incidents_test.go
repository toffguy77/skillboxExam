package prepare

import (
	"github.com/toffguy77/statusPage/internal/models"
	"reflect"
	"testing"
)

func Test_sorted(t *testing.T) {
	type args struct {
		res []models.IncidentData
	}
	tests := []struct {
		name string
		args args
		want []models.IncidentData
	}{
		{
			name: "one active item",
			args: args{
				res: []models.IncidentData{
					{
						Topic:  "Checkout page is down",
						Status: "active",
					},
				},
			},
			want: []models.IncidentData{
				{
					Topic:  "Checkout page is down",
					Status: "active",
				},
			},
		},
		{
			name: "two active items",
			args: args{
				res: []models.IncidentData{
					{
						Topic:  "Checkout page is down",
						Status: "active",
					},
					{
						Topic:  "Support overload",
						Status: "active",
					},
				},
			},
			want: []models.IncidentData{
				{
					Topic:  "Checkout page is down",
					Status: "active",
				},
				{
					Topic:  "Support overload",
					Status: "active",
				},
			},
		},
		{
			name: "two active and one closed items",
			args: args{
				res: []models.IncidentData{
					{
						Topic:  "Checkout page is down",
						Status: "active",
					},
					{
						Topic:  "MMS connection stability",
						Status: "closed",
					},
					{
						Topic:  "Support overload",
						Status: "active",
					},
				},
			},
			want: []models.IncidentData{
				{
					Topic:  "Checkout page is down",
					Status: "active",
				},
				{
					Topic:  "Support overload",
					Status: "active",
				},
				{
					Topic:  "MMS connection stability",
					Status: "closed",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sorted(tt.args.res); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
