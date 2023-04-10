package prepare

import (
	"github.com/toffguy77/statusPage/internal/models"
	"reflect"
	"testing"
)

func Test_changeCountryMMS(t *testing.T) {
	type args struct {
		mmsRes        []models.MMSData
		countriesList map[string]models.Country
	}
	tests := []struct {
		name string
		args args
		want []models.MMSData
	}{
		{
			name: "default",
			args: args{
				mmsRes: []models.MMSData{
					{
						Country:      "AQ",
						Provider:     "Topolo",
						Bandwidth:    "56",
						ResponseTime: "916",
					},
				},
				countriesList: map[string]models.Country{
					"AQ": models.Country{
						Name:       "Антарктика",
						Alpha2:     "AQ",
						Alpha3:     "ATA",
						ISO_3166_1: "10",
						ISO_3166_2: "ISO 3166-2:AQ",
					},
				},
			},
			want: []models.MMSData{
				{
					Country:      "Антарктика",
					Provider:     "Topolo",
					Bandwidth:    "56",
					ResponseTime: "916",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := changeCountryMMS(tt.args.mmsRes, tt.args.countriesList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("changeCountryMMS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortByCountry(t *testing.T) {
	type args struct {
		mmsRes []models.MMSData
	}
	tests := []struct {
		name string
		args args
		want []models.MMSData
	}{
		{
			name: "default",
			args: args{
				mmsRes: []models.MMSData{
					{
						Country:      "Россия",
						Provider:     "Kildy",
						Bandwidth:    "77",
						ResponseTime: "1029",
					},
					{
						Country:      "Антарктика",
						Provider:     "Topolo",
						Bandwidth:    "56",
						ResponseTime: "916",
					},
				},
			},
			want: []models.MMSData{
				{
					Country:      "Антарктика",
					Provider:     "Topolo",
					Bandwidth:    "56",
					ResponseTime: "916",
				},
				{
					Country:      "Россия",
					Provider:     "Kildy",
					Bandwidth:    "77",
					ResponseTime: "1029",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortByCountry(tt.args.mmsRes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortByCountry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortByProvider(t *testing.T) {
	type args struct {
		mmsRes []models.MMSData
	}
	tests := []struct {
		name string
		args args
		want []models.MMSData
	}{
		{
			name: "default",
			args: args{
				mmsRes: []models.MMSData{
					{
						Country:      "Россия",
						Provider:     "Rond",
						Bandwidth:    "69",
						ResponseTime: "776",
					},
					{
						Country:      "Россия",
						Provider:     "Kildy",
						Bandwidth:    "77",
						ResponseTime: "1029",
					},
					{
						Country:      "Антарктика",
						Provider:     "Topolo",
						Bandwidth:    "56",
						ResponseTime: "916",
					},
				},
			},
			want: []models.MMSData{
				{
					Country:      "Россия",
					Provider:     "Kildy",
					Bandwidth:    "77",
					ResponseTime: "1029",
				},
				{
					Country:      "Россия",
					Provider:     "Rond",
					Bandwidth:    "69",
					ResponseTime: "776",
				},
				{
					Country:      "Антарктика",
					Provider:     "Topolo",
					Bandwidth:    "56",
					ResponseTime: "916",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortByProvider(tt.args.mmsRes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortByProvider() = %v, want %v", got, tt.want)
			}
		})
	}
}
