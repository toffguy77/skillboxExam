package countries

import (
	"github.com/toffguy77/statusPage/internal/models"
	"reflect"
	"testing"
)

func Test_isEmptyCountry(t *testing.T) {
	type args struct {
		country models.Country
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "default",
			args: args{country: models.Country{
				Name:       "Another Country",
				Alpha2:     "AA",
				Alpha3:     "AAA",
				ISO_3166_1: "ISO 1111-1:AA",
				ISO_3166_2: "AA",
			}},
			want: false,
		},
		{
			name: "missed name",
			args: args{country: models.Country{
				Name:       "",
				Alpha2:     "AA",
				Alpha3:     "AAA",
				ISO_3166_1: "ISO 1111-1:AA",
				ISO_3166_2: "AA",
			}},
			want: true,
		},
		{
			name: "missed Alpha-2",
			args: args{country: models.Country{
				Name:       "Another Country",
				Alpha2:     "",
				Alpha3:     "AAA",
				ISO_3166_1: "ISO 1111-1:AA",
				ISO_3166_2: "AA",
			}},
			want: true,
		},
		{
			name: "missed Alpha-3",
			args: args{country: models.Country{
				Name:       "Another Country",
				Alpha2:     "AA",
				Alpha3:     "",
				ISO_3166_1: "ISO 1111-1:AA",
				ISO_3166_2: "AA",
			}},
			want: true,
		},
		{
			name: "missed ISO 1",
			args: args{country: models.Country{
				Name:       "Another Country",
				Alpha2:     "AA",
				Alpha3:     "AAA",
				ISO_3166_1: "",
				ISO_3166_2: "AA",
			}},
			want: true,
		},
		{
			name: "missed Alpha-3",
			args: args{country: models.Country{
				Name:       "Another Country",
				Alpha2:     "AA",
				Alpha3:     "AAA",
				ISO_3166_1: "ISO 1111-1:AA",
				ISO_3166_2: "",
			}},
			want: true,
		},
		{
			name: "empty",
			args: args{country: models.Country{
				Name:       "",
				Alpha2:     "",
				Alpha3:     "",
				ISO_3166_1: "",
				ISO_3166_2: "",
			}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEmptyCountry(tt.args.country); got != tt.want {
				t.Errorf("isEmptyCountry() = %v, want %v", got, tt.want)
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
		want models.Country
	}{
		{
			name: "default",
			args: args{line: "Антарктика;AQ;ATA;10;ISO 3166-2:AQ"},
			want: models.Country{
				Name:       "Антарктика",
				Alpha2:     "AQ",
				Alpha3:     "ATA",
				ISO_3166_1: "10",
				ISO_3166_2: "ISO 3166-2:AQ",
			},
		},
		{
			name: "missed field",
			args: args{line: "Антарктика;AQ;10;ISO 3166-2:AQ"},
			want: models.Country{},
		},
		{
			name: "empty structured field",
			args: args{line: "Антарктика;AQ;;10;ISO 3166-2:AQ"},
			want: models.Country{},
		},
		{
			name: "empty non structured field",
			args: args{line: "Антарктика;AQ;ATA;10;"},
			want: models.Country{
				Name:       "Антарктика",
				Alpha2:     "AQ",
				Alpha3:     "ATA",
				ISO_3166_1: "10",
				ISO_3166_2: "",
			},
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
