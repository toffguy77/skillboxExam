package countries

import (
	"github.com/toffguy77/statusPage/internal/models"
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEmptyCountry(tt.args.country); got != tt.want {
				t.Errorf("isEmptyCountry() = %v, want %v", got, tt.want)
			}
		})
	}
}
