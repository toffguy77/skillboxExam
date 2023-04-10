package voicecall

import (
	"github.com/toffguy77/statusPage/internal/models"
	"testing"
)

func Test_isEmptyVoiceData(t *testing.T) {
	type args struct {
		data models.VoiceCallData
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
			if got := isEmptyVoiceData(tt.args.data); got != tt.want {
				t.Errorf("isEmptyVoiceData() = %v, want %v", got, tt.want)
			}
		})
	}
}
