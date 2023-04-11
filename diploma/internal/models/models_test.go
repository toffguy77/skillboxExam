package models

import (
	"testing"
)

func TestMMSData_HasCountry(t *testing.T) {
	type fields struct {
		Country      string
		Provider     string
		Bandwidth    string
		ResponseTime string
	}
	type args struct {
		country map[string]Country
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "default",
			fields: fields{
				Country:      "RU",
				Provider:     "Topolo",
				Bandwidth:    "85",
				ResponseTime: "911",
			},
			args: args{country: map[string]Country{
				"GB": {
					Name:       "Великобритания",
					Alpha2:     "GB",
					Alpha3:     "GBR",
					ISO_3166_1: "826",
					ISO_3166_2: "ISO 3166-2:GB",
				},
				"RU": {
					Name:       "Россия",
					Alpha2:     "RU",
					Alpha3:     "RUS",
					ISO_3166_1: "643",
					ISO_3166_2: "ISO 3166-2:RU",
				},
			}},
			want: true,
		},
		{
			name: "missed",
			fields: fields{
				Country:      "RU",
				Provider:     "Topolo",
				Bandwidth:    "85",
				ResponseTime: "911",
			},
			args: args{country: map[string]Country{
				"GB": Country{
					Name:       "Великобритания",
					Alpha2:     "GB",
					Alpha3:     "GBR",
					ISO_3166_1: "826",
					ISO_3166_2: "ISO 3166-2:GB",
				},
			}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MMSData{
				Country:      tt.fields.Country,
				Provider:     tt.fields.Provider,
				Bandwidth:    tt.fields.Bandwidth,
				ResponseTime: tt.fields.ResponseTime,
			}
			if got := m.HasCountry(tt.args.country); got != tt.want {
				t.Errorf("HasCountry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmailData_HasCountry(t *testing.T) {
	type fields struct {
		Country      string
		Provider     string
		DeliveryTime int
	}
	type args struct {
		country map[string]Country
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "default",
			fields: fields{
				Country:      "RU",
				Provider:     "Hotmail",
				DeliveryTime: 487,
			},
			args: args{country: map[string]Country{
				"GB": {
					Name:       "Великобритания",
					Alpha2:     "GB",
					Alpha3:     "GBR",
					ISO_3166_1: "826",
					ISO_3166_2: "ISO 3166-2:GB",
				},
				"RU": {
					Name:       "Россия",
					Alpha2:     "RU",
					Alpha3:     "RUS",
					ISO_3166_1: "643",
					ISO_3166_2: "ISO 3166-2:RU",
				},
			}},
			want: true,
		},
		{
			name: "missed",
			fields: fields{
				Country:      "RU",
				Provider:     "Hotmail",
				DeliveryTime: 487,
			},
			args: args{country: map[string]Country{
				"GB": Country{
					Name:       "Великобритания",
					Alpha2:     "GB",
					Alpha3:     "GBR",
					ISO_3166_1: "826",
					ISO_3166_2: "ISO 3166-2:GB",
				},
			}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := EmailData{
				Country:      tt.fields.Country,
				Provider:     tt.fields.Provider,
				DeliveryTime: tt.fields.DeliveryTime,
			}
			if got := m.HasCountry(tt.args.country); got != tt.want {
				t.Errorf("HasCountry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSMSData_HasCountry(t *testing.T) {
	type fields struct {
		Country      string
		Bandwidth    string
		ResponseTime string
		Provider     string
	}
	type args struct {
		country map[string]Country
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "default",
			fields: fields{
				Country:      "RU",
				Bandwidth:    "1",
				ResponseTime: "499",
				Provider:     "Topolo",
			},
			args: args{country: map[string]Country{
				"GB": {
					Name:       "Великобритания",
					Alpha2:     "GB",
					Alpha3:     "GBR",
					ISO_3166_1: "826",
					ISO_3166_2: "ISO 3166-2:GB",
				},
				"RU": {
					Name:       "Россия",
					Alpha2:     "RU",
					Alpha3:     "RUS",
					ISO_3166_1: "643",
					ISO_3166_2: "ISO 3166-2:RU",
				},
			}},
			want: true,
		},
		{
			name: "missed",
			fields: fields{
				Country:      "RU",
				Bandwidth:    "1",
				ResponseTime: "499",
				Provider:     "Topolo",
			},
			args: args{country: map[string]Country{
				"GB": Country{
					Name:       "Великобритания",
					Alpha2:     "GB",
					Alpha3:     "GBR",
					ISO_3166_1: "826",
					ISO_3166_2: "ISO 3166-2:GB",
				},
			}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := SMSData{
				Country:      tt.fields.Country,
				Bandwidth:    tt.fields.Bandwidth,
				ResponseTime: tt.fields.ResponseTime,
				Provider:     tt.fields.Provider,
			}
			if got := m.HasCountry(tt.args.country); got != tt.want {
				t.Errorf("HasCountry() = %v, want %v", got, tt.want)
			}
		})
	}
}
