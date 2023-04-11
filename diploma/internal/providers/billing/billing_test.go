package billing

import (
	"github.com/toffguy77/statusPage/internal/models"
	"reflect"
	"testing"
)

func Test_calcNumFromMask(t *testing.T) {
	type args struct {
		mask string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "default",
			args: args{mask: "010011"},
			want: 19,
		},
		{
			name: "zeros",
			args: args{mask: "000000"},
			want: 0,
		},
		{
			name: "ones",
			args: args{mask: "111111"},
			want: 63,
		},
		{
			name: "bigger mask",
			args: args{mask: "1111111"},
			want: 127,
		},
		{
			name: "fewer mask",
			args: args{mask: "11111"},
			want: 31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcNumFromMask(tt.args.mask); got != tt.want {
				t.Errorf("calcNumFromMask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseMask(t *testing.T) {
	type args struct {
		mask string
	}
	tests := []struct {
		name string
		args args
		want models.BillingData
	}{
		{
			name: "default",
			args: args{mask: "010011"},
			want: models.BillingData{
				CreateCustomer: true,
				Purchase:       true,
				Payout:         false,
				Recurring:      false,
				FraudControl:   true,
				CheckoutPage:   false,
			},
		},
		{
			name: "zeros",
			args: args{mask: "000000"},
			want: models.BillingData{
				CreateCustomer: false,
				Purchase:       false,
				Payout:         false,
				Recurring:      false,
				FraudControl:   false,
				CheckoutPage:   false,
			},
		},
		{
			name: "ones",
			args: args{mask: "111111"},
			want: models.BillingData{
				CreateCustomer: true,
				Purchase:       true,
				Payout:         true,
				Recurring:      true,
				FraudControl:   true,
				CheckoutPage:   true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseMask(tt.args.mask); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseMask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validate(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "default",
			args:    args{data: "010011"},
			wantErr: false,
		},
		{
			name:    "fewer digits",
			args:    args{data: "10011"},
			wantErr: true,
		},
		{
			name:    "more digits",
			args:    args{data: "1010011"},
			wantErr: true,
		},
		{
			name:    "more than just ones and zeros",
			args:    args{data: "010a011"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validate(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
