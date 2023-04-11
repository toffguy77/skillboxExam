package prepare

import (
	"github.com/toffguy77/statusPage/internal/models"
	"reflect"
	"testing"
)

func Test_changeCountrySMS(t *testing.T) {
	type args struct {
		smsRes        []models.SMSData
		countriesList map[string]models.Country
	}
	tests := []struct {
		name string
		args args
		want []models.SMSData
	}{
		{
			name: "default",
			args: args{
				smsRes: []models.SMSData{
					{
						Country:      "AQ",
						Bandwidth:    "93",
						ResponseTime: "1433",
						Provider:     "Rond",
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
			want: []models.SMSData{
				{
					Country:      "Антарктика",
					Bandwidth:    "93",
					ResponseTime: "1433",
					Provider:     "Rond",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := changeCountrySMS(tt.args.smsRes, tt.args.countriesList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("changeCountrySMS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortSmsByCountry(t *testing.T) {
	type args struct {
		smsRes []models.SMSData
	}
	tests := []struct {
		name string
		args args
		want []models.SMSData
	}{
		{
			name: "default",
			args: args{
				smsRes: []models.SMSData{
					{
						Country:      "Россия",
						Bandwidth:    "77",
						ResponseTime: "1024",
						Provider:     "Kildy",
					},
					{
						Country:      "Антарктика",
						Bandwidth:    "93",
						ResponseTime: "1433",
						Provider:     "Rond",
					},
				},
			},
			want: []models.SMSData{
				{
					Country:      "Антарктика",
					Bandwidth:    "93",
					ResponseTime: "1433",
					Provider:     "Rond",
				},
				{
					Country:      "Россия",
					Bandwidth:    "77",
					ResponseTime: "1024",
					Provider:     "Kildy",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortSmsByCountry(tt.args.smsRes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortSmsByCountry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortSmsByProvider(t *testing.T) {
	type args struct {
		smsRes []models.SMSData
	}
	tests := []struct {
		name string
		args args
		want []models.SMSData
	}{
		{
			name: "default",
			args: args{
				smsRes: []models.SMSData{
					{
						Country:      "Антарктика",
						Bandwidth:    "93",
						ResponseTime: "1433",
						Provider:     "Rond",
					},
					{
						Country:      "Россия",
						Bandwidth:    "77",
						ResponseTime: "1024",
						Provider:     "Kildy",
					},
				},
			},
			want: []models.SMSData{
				{
					Country:      "Россия",
					Bandwidth:    "77",
					ResponseTime: "1024",
					Provider:     "Kildy",
				},
				{
					Country:      "Антарктика",
					Bandwidth:    "93",
					ResponseTime: "1433",
					Provider:     "Rond",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortSmsByProvider(tt.args.smsRes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortSmsByProvider() = %v, want %v", got, tt.want)
			}
		})
	}
}
