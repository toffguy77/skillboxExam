package prepare

import (
	"github.com/toffguy77/statusPage/internal/models"
	"reflect"
	"testing"
)

func Test_resByCountry(t *testing.T) {
	type args struct {
		testData []models.EmailData
	}
	tests := []struct {
		name string
		args args
		want map[string][]models.EmailData
	}{
		{
			name: "default",
			args: args{testData: []models.EmailData{{
				Country:      "RU",
				Provider:     "Yahoo",
				DeliveryTime: 554,
			},
				{
					Country:      "EN",
					Provider:     "Hotmail",
					DeliveryTime: 487,
				},
			}},
			want: map[string][]models.EmailData{
				"EN": []models.EmailData{{
					Country:      "EN",
					Provider:     "Hotmail",
					DeliveryTime: 487,
				}},
				"RU": []models.EmailData{{
					Country:      "RU",
					Provider:     "Yahoo",
					DeliveryTime: 554,
				}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resByCountry(tt.args.testData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resByCountry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortedByDeliveryTime(t *testing.T) {
	type args struct {
		emailResByCountry map[string][]models.EmailData
	}
	tests := []struct {
		name string
		args args
		want map[string][][]models.EmailData
	}{
		{
			name: "one item in the list",
			args: args{emailResByCountry: map[string][]models.EmailData{
				"RU": []models.EmailData{
					{
						Country:      "RU",
						Provider:     "Hotmail",
						DeliveryTime: 487,
					},
				},
				"EN": []models.EmailData{
					{
						Country:      "EN",
						Provider:     "Yahoo",
						DeliveryTime: 554,
					},
				},
			},
			},
			want: map[string][][]models.EmailData{
				"RU": [][]models.EmailData{
					{
						{
							Country:      "RU",
							Provider:     "Hotmail",
							DeliveryTime: 487,
						},
					},
					{
						{
							Country:      "RU",
							Provider:     "Hotmail",
							DeliveryTime: 487,
						}}},
				"EN": [][]models.EmailData{
					{
						{
							Country:      "EN",
							Provider:     "Yahoo",
							DeliveryTime: 554,
						},
					},
					{
						{
							Country:      "EN",
							Provider:     "Yahoo",
							DeliveryTime: 554,
						},
					},
				},
			},
		},
		{
			name: "two items in the list",
			args: args{emailResByCountry: map[string][]models.EmailData{
				"RU": []models.EmailData{
					{
						Country:      "RU",
						Provider:     "Hotmail",
						DeliveryTime: 487,
					},
					{
						Country:      "RU",
						Provider:     "Yahoo",
						DeliveryTime: 554,
					},
				},
				"EN": []models.EmailData{
					{
						Country:      "EN",
						Provider:     "Yahoo",
						DeliveryTime: 554,
					},
					{
						Country:      "EN",
						Provider:     "Hotmail",
						DeliveryTime: 487,
					},
				},
			},
			},
			want: map[string][][]models.EmailData{
				"RU": [][]models.EmailData{
					{
						{
							Country:      "RU",
							Provider:     "Hotmail",
							DeliveryTime: 487,
						},
						{
							Country:      "RU",
							Provider:     "Yahoo",
							DeliveryTime: 554,
						},
					},
					{
						{
							Country:      "RU",
							Provider:     "Yahoo",
							DeliveryTime: 554,
						},
						{
							Country:      "RU",
							Provider:     "Hotmail",
							DeliveryTime: 487,
						}}},
				"EN": [][]models.EmailData{
					{
						{
							Country:      "EN",
							Provider:     "Hotmail",
							DeliveryTime: 487,
						},
						{
							Country:      "EN",
							Provider:     "Yahoo",
							DeliveryTime: 554,
						},
					},
					{
						{
							Country:      "EN",
							Provider:     "Yahoo",
							DeliveryTime: 554,
						},
						{
							Country:      "EN",
							Provider:     "Hotmail",
							DeliveryTime: 487,
						},
					},
				},
			},
		},
		{
			name: "four items in the list",
			args: args{emailResByCountry: map[string][]models.EmailData{
				"RU": []models.EmailData{
					{
						Country:      "RU",
						Provider:     "Hotmail",
						DeliveryTime: 487,
					},
					{
						Country:      "RU",
						Provider:     "Yahoo",
						DeliveryTime: 554,
					},
					{
						Country:      "RU",
						Provider:     "Hotmail",
						DeliveryTime: 113,
					},
					{
						Country:      "RU",
						Provider:     "Yahoo",
						DeliveryTime: 555,
					},
				},
				"EN": []models.EmailData{
					{
						Country:      "EN",
						Provider:     "Hotmail",
						DeliveryTime: 113,
					},
					{
						Country:      "EN",
						Provider:     "Yahoo",
						DeliveryTime: 554,
					},
					{
						Country:      "EN",
						Provider:     "Hotmail",
						DeliveryTime: 487,
					},
					{
						Country:      "EN",
						Provider:     "Hotmail",
						DeliveryTime: 555,
					},
				},
			},
			},
			want: map[string][][]models.EmailData{
				"RU": [][]models.EmailData{
					{
						{
							Country:      "RU",
							Provider:     "Hotmail",
							DeliveryTime: 113,
						},
						{
							Country:      "RU",
							Provider:     "Hotmail",
							DeliveryTime: 487,
						},
						{
							Country:      "RU",
							Provider:     "Yahoo",
							DeliveryTime: 554,
						},
					},
					{
						{
							Country:      "RU",
							Provider:     "Yahoo",
							DeliveryTime: 555,
						},
						{
							Country:      "RU",
							Provider:     "Yahoo",
							DeliveryTime: 554,
						},
						{
							Country:      "RU",
							Provider:     "Hotmail",
							DeliveryTime: 487,
						}}},
				"EN": [][]models.EmailData{
					{
						{
							Country:      "EN",
							Provider:     "Hotmail",
							DeliveryTime: 113,
						},
						{
							Country:      "EN",
							Provider:     "Hotmail",
							DeliveryTime: 487,
						},
						{
							Country:      "EN",
							Provider:     "Yahoo",
							DeliveryTime: 554,
						},
					},
					{
						{
							Country:      "EN",
							Provider:     "Hotmail",
							DeliveryTime: 555,
						},
						{
							Country:      "EN",
							Provider:     "Yahoo",
							DeliveryTime: 554,
						},
						{
							Country:      "EN",
							Provider:     "Hotmail",
							DeliveryTime: 487,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedByDeliveryTime(tt.args.emailResByCountry); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedByDeliveryTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortByDeliveryAsc(t *testing.T) {
	type args struct {
		data []models.EmailData
	}
	tests := []struct {
		name string
		args args
		want []models.EmailData
	}{
		{
			name: "one items in the list",
			args: args{data: []models.EmailData{
				{
					Country:      "RU",
					Provider:     "Yahoo",
					DeliveryTime: 554,
				},
			}},
			want: []models.EmailData{
				{
					Country:      "RU",
					Provider:     "Yahoo",
					DeliveryTime: 554,
				},
			},
		},
		{
			name: "two items in the list",
			args: args{data: []models.EmailData{
				{
					Country:      "RU",
					Provider:     "Hotmail",
					DeliveryTime: 487,
				},
				{
					Country:      "RU",
					Provider:     "Yahoo",
					DeliveryTime: 555,
				},
			}},
			want: []models.EmailData{
				{
					Country:      "RU",
					Provider:     "Hotmail",
					DeliveryTime: 487,
				},
				{
					Country:      "RU",
					Provider:     "Yahoo",
					DeliveryTime: 555,
				},
			},
		},
		{
			name: "four items in the list",
			args: args{data: []models.EmailData{
				{
					Country:      "RU",
					Provider:     "Hotmail",
					DeliveryTime: 487,
				},
				{
					Country:      "RU",
					Provider:     "Yahoo",
					DeliveryTime: 554,
				},
				{
					Country:      "RU",
					Provider:     "Hotmail",
					DeliveryTime: 113,
				},
				{
					Country:      "RU",
					Provider:     "Yahoo",
					DeliveryTime: 555,
				},
			}},
			want: []models.EmailData{
				{
					Country:      "RU",
					Provider:     "Hotmail",
					DeliveryTime: 113,
				},
				{
					Country:      "RU",
					Provider:     "Hotmail",
					DeliveryTime: 487,
				},
				{
					Country:      "RU",
					Provider:     "Yahoo",
					DeliveryTime: 554,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortByDeliveryAsc(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortByDeliveryAsc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortByDeliveryDesc(t *testing.T) {
	type args struct {
		data []models.EmailData
	}
	tests := []struct {
		name string
		args args
		want []models.EmailData
	}{
		{
			name: "one items in the list",
			args: args{data: []models.EmailData{
				{
					Country:      "RU",
					Provider:     "Yahoo",
					DeliveryTime: 554,
				},
			}},
			want: []models.EmailData{
				{
					Country:      "RU",
					Provider:     "Yahoo",
					DeliveryTime: 554,
				},
			},
		},
		{
			name: "two items in the list",
			args: args{data: []models.EmailData{
				{
					Country:      "RU",
					Provider:     "Yahoo",
					DeliveryTime: 555,
				},
				{
					Country:      "RU",
					Provider:     "Hotmail",
					DeliveryTime: 487,
				},
			}},
			want: []models.EmailData{
				{
					Country:      "RU",
					Provider:     "Yahoo",
					DeliveryTime: 555,
				},
				{
					Country:      "RU",
					Provider:     "Hotmail",
					DeliveryTime: 487,
				},
			},
		},
		{
			name: "four items in the list",
			args: args{data: []models.EmailData{
				{
					Country:      "RU",
					Provider:     "Hotmail",
					DeliveryTime: 487,
				},
				{
					Country:      "RU",
					Provider:     "Yahoo",
					DeliveryTime: 554,
				},
				{
					Country:      "RU",
					Provider:     "Hotmail",
					DeliveryTime: 113,
				},
				{
					Country:      "RU",
					Provider:     "Yahoo",
					DeliveryTime: 555,
				},
			}},
			want: []models.EmailData{
				{
					Country:      "RU",
					Provider:     "Yahoo",
					DeliveryTime: 555,
				},
				{
					Country:      "RU",
					Provider:     "Yahoo",
					DeliveryTime: 554,
				},
				{
					Country:      "RU",
					Provider:     "Hotmail",
					DeliveryTime: 487,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortByDeliveryDesc(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortByDeliveryDesc() = %v, want %v", got, tt.want)
			}
		})
	}
}
