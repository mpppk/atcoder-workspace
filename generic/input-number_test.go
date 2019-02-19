package generic

import (
	"reflect"
	"testing"
)

func TestInput_GetAAALine(t *testing.T) {
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
		want    []AAA
		wantErr bool
	}{
		{
			name: "GetCValueLine",
			fields: fields{
				lines: [][]string{{"1", "2", "3"}, {"4", "5", "6"}},
			},
			args: args{
				index: 0,
			},
			want:    []AAA{1, 2, 3},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Input{
				lines: tt.fields.lines,
			}
			got, err := i.GetAAALine(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("Input.GetValueLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Input.GetValueLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
