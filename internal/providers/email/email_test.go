package email

import (
	"github.com/toffguy77/statusPage/internal/models"
	"reflect"
	"testing"
)

func TestEmailProvider_GetStatus(t *testing.T) {
	type fields struct {
		Name string
	}
	type args struct {
		countries map[string]models.Country
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.EmailData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := EmailProvider{
				Name: tt.fields.Name,
			}
			got, err := p.GetStatus(tt.args.countries)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStatus() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isTrustedEmailProvider(t *testing.T) {
	type args struct {
		provider string
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
			if got := isTrustedEmailProvider(tt.args.provider); got != tt.want {
				t.Errorf("isTrustedEmailProvider() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseEmailData(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		want    []models.EmailData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseEmailData(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseEmailData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseEmailData() got = %v, want %v", got, tt.want)
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
		want models.EmailData
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseString(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseString() = %v, want %v", got, tt.want)
			}
		})
	}
}
