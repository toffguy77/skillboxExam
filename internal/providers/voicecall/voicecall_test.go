package voicecall

import (
	"github.com/toffguy77/statusPage/internal/models"
	"reflect"
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
		{
			name: "default",
			args: args{data: models.VoiceCallData{
				Country:             "RU",
				Bandwidth:           "44",
				ResponseTime:        "669",
				Provider:            "TransparentCalls",
				ConnectionStability: 0.71,
				TTFB:                904,
				VoicePurity:         26,
				MedianOfCallsTime:   39,
			}},
			want: false,
		},
		{
			name: "empty",
			args: args{data: models.VoiceCallData{}},
			want: true,
		},
		{
			name: "some fields missed",
			args: args{data: models.VoiceCallData{
				Country:             "RU",
				Bandwidth:           "44",
				ResponseTime:        "",
				Provider:            "TransparentCalls",
				ConnectionStability: 0,
				TTFB:                904,
				VoicePurity:         26,
				MedianOfCallsTime:   39,
			}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEmptyVoiceData(tt.args.data); got != tt.want {
				t.Errorf("isEmptyVoiceData() = %v, want %v", got, tt.want)
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
		want models.VoiceCallData
	}{
		{
			name: "default",
			args: args{line: "RU;44;669;TransparentCalls;0.71;904;26;39"},
			want: models.VoiceCallData{
				Country:             "RU",
				Bandwidth:           "44",
				ResponseTime:        "669",
				Provider:            "TransparentCalls",
				ConnectionStability: 0.71,
				TTFB:                904,
				VoicePurity:         26,
				MedianOfCallsTime:   39,
			},
		},
		{
			name: "more fields",
			args: args{line: "RU;44;669;TransparentCalls;0.71;904;26;39;42"},
			want: models.VoiceCallData{},
		},
		{
			name: "less fields",
			args: args{line: "RU;44;669;TransparentCalls;0.71;904;26"},
			want: models.VoiceCallData{},
		},
		{
			name: "provider is not trusted",
			args: args{line: "RU;44;669;Rambler;0.71;904;26;39"},
			want: models.VoiceCallData{},
		},
		{
			name: "ConnectionStability is not float",
			args: args{line: "RU;44;669;TransparentCalls;0;90a;26;39"},
			want: models.VoiceCallData{},
		},
		{
			name: "ttfb is not integer",
			args: args{line: "RU;44;669;TransparentCalls;0.71;904.0;26;39"},
			want: models.VoiceCallData{},
		},
		{
			name: "VoicePurity is not integer",
			args: args{line: "RU;44;669;TransparentCalls;0.71;904;aa;39"},
			want: models.VoiceCallData{},
		},
		{
			name: "medianOfCallsTime is not integer",
			args: args{line: "RU;44;669;TransparentCalls;0.71;904;26;39.0"},
			want: models.VoiceCallData{},
		},
		{
			name: "default",
			args: args{line: "RU;44;669;TransparentCalls;0.71;904;26;39"},
			want: models.VoiceCallData{
				Country:             "RU",
				Bandwidth:           "44",
				ResponseTime:        "669",
				Provider:            "TransparentCalls",
				ConnectionStability: 0.71,
				TTFB:                904,
				VoicePurity:         26,
				MedianOfCallsTime:   39,
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

func Test_isCorrectLine(t *testing.T) {
	type args struct {
		line []string
		l    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "default",
			args: args{
				line: []string{"nil"},
				l:    1,
			},
			want: true,
		},
		{
			name: "default x3",
			args: args{
				line: []string{"nil", "zero", "empty"},
				l:    3,
			},
			want: true,
		},
		{
			name: "default",
			args: args{
				line: []string{},
				l:    1,
			},
			want: false,
		},
		{
			name: "not correct",
			args: args{
				line: []string{"nil", "zero", "empty"},
				l:    2,
			},
			want: false,
		},
		{
			name: "with empty item",
			args: args{
				line: []string{"nil", "", "empty"},
				l:    3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCorrectLine(tt.args.line, tt.args.l); got != tt.want {
				t.Errorf("isCorrectLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
