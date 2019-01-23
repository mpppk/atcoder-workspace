package utl

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func Test_toLines(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  [][]string
	}{
		{
			name: "can convert scanner to [][]string",
			input: `
					1
					2 3

					test hoge fuga
			`,
			want: [][]string{
				{"1"},
				{"2", "3"},
				{},
				{"test", "hoge", "fuga"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(TrimSpaceAndNewLineCodeAndTab(tt.input)))
			if got := toLines(scanner); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toIntLine(t *testing.T) {
	type args struct {
		line []string
	}
	tests := []struct {
		name        string
		args        args
		wantIntLine []int64
		wantErr     bool
	}{
		{
			name: "can convert from string slice to int64 slice",
			args: args{
				line: []string{"1", "2", "3"},
			},
			wantIntLine: []int64{1, 2, 3},
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIntLine, err := toInt64Line(tt.args.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("toSpecificBitIntLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotIntLine, tt.wantIntLine) {
				t.Errorf("toSpecificBitIntLine() = %v, want %v", gotIntLine, tt.wantIntLine)
			}
		})
	}
}

func TestInput_GetLine(t *testing.T) {
	type fields struct {
		lines [][]string
	}
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "can get specified index line",
			fields: fields{
				lines: [][]string{
					{"1", "2", "3"},
					{"4", "5", "6"},
					{"7", "8", "9"},
				},
			},
			args:    args{index: 0},
			want:    []string{"1", "2", "3"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Input{
				lines: tt.fields.lines,
			}
			got, err := i.GetLine(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("Input.GetLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Input.GetLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
