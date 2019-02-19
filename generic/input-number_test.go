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

func TestInput_GetAAAValue(t *testing.T) {
	type fields struct {
		lines [][]string
	}
	type args struct {
		rowIndex int
		colIndex int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    AAA
		wantErr bool
	}{
		{
			name: "GetAAAValue",
			fields: fields{
				lines: [][]string{{"1", "2", "3"}, {"4", "5", "6"}},
			},
			args: args{
				rowIndex: 0,
				colIndex: 0,
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "GetAAAValue",
			fields: fields{
				lines: [][]string{{"1", "2", "3"}, {"4", "5", "6"}},
			},
			args: args{
				rowIndex: 1,
				colIndex: 2,
			},
			want:    6,
			wantErr: false,
		},
		{
			name: "GetAAAValue",
			fields: fields{
				lines: [][]string{{"1", "2", "3"}, {"4", "5", "6"}},
			},
			args: args{
				rowIndex: 2,
				colIndex: 2,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "GetAAAValue",
			fields: fields{
				lines: [][]string{{"1", "2", "3"}, {"4", "5", "6"}},
			},
			args: args{
				rowIndex: 1,
				colIndex: 3,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Input{
				lines: tt.fields.lines,
			}
			got, err := i.GetAAAValue(tt.args.rowIndex, tt.args.colIndex)
			if (err != nil) != tt.wantErr {
				t.Errorf("Input.GetAAAValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Input.GetAAAValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInput_GetFirstAAAValue(t *testing.T) {
	type fields struct {
		lines [][]string
	}
	type args struct {
		rowIndex int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    AAA
		wantErr bool
	}{
		{
			name: "GetFirstAAAValue",
			fields: fields{
				lines: [][]string{{"1", "2", "3"}, {"4", "5", "6"}},
			},
			args: args{
				rowIndex: 0,
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "GetFirstAAAValue",
			fields: fields{
				lines: [][]string{{"1", "2", "3"}, {"4", "5", "6"}},
			},
			args: args{
				rowIndex: 1,
			},
			want:    4,
			wantErr: false,
		},
		{
			name: "GetFirstAAAValue",
			fields: fields{
				lines: [][]string{{"1", "2", "3"}, {"4", "5", "6"}},
			},
			args: args{
				rowIndex: 2,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Input{
				lines: tt.fields.lines,
			}
			got, err := i.GetFirstAAAValue(tt.args.rowIndex)
			if (err != nil) != tt.wantErr {
				t.Errorf("Input.GetFirstAAAValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Input.GetFirstAAAValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInput_GetColAAALine(t *testing.T) {
	type fields struct {
		lines [][]string
	}
	type args struct {
		colIndex int
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantNewLine []AAA
		wantErr     bool
	}{
		{
			name: "GetColAAALine",
			fields: fields{
				lines: [][]string{{"1", "2", "3"}, {"4", "5", "6"}},
			},
			args: args{
				colIndex: 0,
			},
			wantNewLine: []AAA{1, 4},
			wantErr:     false,
		},
		{
			name: "GetColAAALine",
			fields: fields{
				lines: [][]string{{"1", "2", "3"}, {"4", "5", "6"}},
			},
			args: args{
				colIndex: 2,
			},
			wantNewLine: []AAA{3, 6},
			wantErr:     false,
		},
		{
			name: "GetColAAALine will throw error if out of range col index is given",
			fields: fields{
				lines: [][]string{{"1", "2", "3"}, {"4", "5", "6"}},
			},
			args: args{
				colIndex: 3,
			},
			wantNewLine: nil,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Input{
				lines: tt.fields.lines,
			}
			gotNewLine, err := i.GetColAAALine(tt.args.colIndex)
			if (err != nil) != tt.wantErr {
				t.Errorf("Input.GetColAAALine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotNewLine, tt.wantNewLine) {
				t.Errorf("Input.GetColAAALine() = %v, want %v", gotNewLine, tt.wantNewLine)
			}
		})
	}
}

func TestInput_GetAAALines(t *testing.T) {
	type fields struct {
		lines [][]string
	}
	tests := []struct {
		name         string
		fields       fields
		wantNewLines [][]AAA
		wantErr      bool
	}{
		{
			name: "GetAAALines",
			fields: fields{
				lines: [][]string{{"1", "2", "3"}, {"4", "5", "6"}},
			},
			wantNewLines: [][]AAA{{1, 2, 3}, {4, 5, 6}},
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Input{
				lines: tt.fields.lines,
			}
			gotNewLines, err := i.GetAAALines()
			if (err != nil) != tt.wantErr {
				t.Errorf("Input.GetAAALines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotNewLines, tt.wantNewLines) {
				t.Errorf("Input.GetAAALines() = %v, want %v", gotNewLines, tt.wantNewLines)
			}
		})
	}
}
