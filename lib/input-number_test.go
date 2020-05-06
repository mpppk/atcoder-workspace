package lib

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

func TestInput_GetFirstAndSecondAAAValue(t *testing.T) {
	type fields struct {
		lines [][]string
	}
	type args struct {
		rowIndex int
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantFirst  AAA
		wantSecond AAA
		wantErr    bool
	}{
		{
			name: "GetFirstAndSecondAAAValue",
			fields: fields{
				lines: [][]string{{"1", "2", "3"}, {"4", "5", "6"}},
			},
			args: args{
				rowIndex: 0,
			},
			wantFirst:  1,
			wantSecond: 2,
			wantErr:    false,
		},
		{
			name: "GetFirstAndSecondAAAValue",
			fields: fields{
				lines: [][]string{{"1", "2", "3"}, {"4", "5", "6"}},
			},
			args: args{
				rowIndex: 1,
			},
			wantFirst:  4,
			wantSecond: 5,
			wantErr:    false,
		},
		{
			name: "GetFirstAndSecondAAAValue",
			fields: fields{
				lines: [][]string{{"1", "2", "3"}, {"4", "5", "6"}},
			},
			args: args{
				rowIndex: 2,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Input{
				lines: tt.fields.lines,
			}
			first, second, err := i.GetFirstAndSecondAAAValue(tt.args.rowIndex)
			if (err != nil) != tt.wantErr {
				t.Errorf("Input.GetFirstAAAValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(first, tt.wantFirst) {
				t.Errorf("Input.GetFirstAAAValue() = %v, want %v", first, tt.wantFirst)
			}
			if !reflect.DeepEqual(second, tt.wantSecond) {
				t.Errorf("Input.GetFirstAAAValue() = %v, want %v", second, tt.wantSecond)
			}
		})
	}
}

func TestInput_GetFromFirstToThirdAAAValue(t *testing.T) {
	type fields struct {
		lines [][]string
	}
	type args struct {
		rowIndex int
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantFirst  AAA
		wantSecond AAA
		wantThird  AAA
		wantErr    bool
	}{
		{
			name: "GetFirstAndSecondAAAValue",
			fields: fields{
				lines: [][]string{{"1", "2", "3"}, {"4", "5", "6"}},
			},
			args: args{
				rowIndex: 0,
			},
			wantFirst:  1,
			wantSecond: 2,
			wantThird:  3,
			wantErr:    false,
		},
		{
			name: "GetFirstAndSecondAAAValue",
			fields: fields{
				lines: [][]string{{"1", "2", "3"}, {"4", "5", "6"}},
			},
			args: args{
				rowIndex: 1,
			},
			wantFirst:  4,
			wantSecond: 5,
			wantThird:  6,
			wantErr:    false,
		},
		{
			name: "GetFirstAndSecondAAAValue",
			fields: fields{
				lines: [][]string{{"1", "2", "3"}, {"4", "5", "6"}},
			},
			args: args{
				rowIndex: 2,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Input{
				lines: tt.fields.lines,
			}
			first, second, third, err := i.GetFromFirstToThirdAAAValue(tt.args.rowIndex)
			if (err != nil) != tt.wantErr {
				t.Errorf("Input.GetFirstAAAValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(first, tt.wantFirst) {
				t.Errorf("Input.GetFirstAAAValue() = %v, want %v", first, tt.wantFirst)
			}
			if !reflect.DeepEqual(second, tt.wantSecond) {
				t.Errorf("Input.GetFirstAAAValue() = %v, want %v", second, tt.wantSecond)
			}
			if !reflect.DeepEqual(third, tt.wantThird) {
				t.Errorf("Input.GetFirstAAAValue() = %v, want %v", third, tt.wantThird)
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

func TestInput_GetAAALinesFrom(t *testing.T) {
	type fields struct {
		lines [][]string
	}
	type args struct {
		fromIndex int
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantNewLines [][]AAA
		wantErr      bool
	}{
		{
			name: "GetAAALines",
			fields: fields{
				lines: [][]string{{"1", "2", "3"}, {"4", "5", "6"}},
			},
			args: args{
				fromIndex: 0,
			},
			wantNewLines: [][]AAA{{1, 2, 3}, {4, 5, 6}},
			wantErr:      false,
		},
		{
			name: "GetAAALines",
			fields: fields{
				lines: [][]string{{"1", "2", "3"}, {"4", "5", "6"}},
			},
			args: args{
				fromIndex: 1,
			},
			wantNewLines: [][]AAA{{4, 5, 6}},
			wantErr:      false,
		},
		{
			name: "GetAAALines",
			fields: fields{
				lines: [][]string{{"1", "2", "3"}, {"4", "5", "6"}},
			},
			args: args{
				fromIndex: 2,
			},
			wantNewLines: nil,
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Input{
				lines: tt.fields.lines,
			}
			gotNewLines, err := i.GetAAALinesFrom(tt.args.fromIndex)
			if (err != nil) != tt.wantErr {
				t.Errorf("Input.GetAAALinesFrom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotNewLines, tt.wantNewLines) {
				t.Errorf("Input.GetAAALinesFrom() = %v, want %v", gotNewLines, tt.wantNewLines)
			}
		})
	}
}

func TestInput_GetAAALineRange(t *testing.T) {
	type fields struct {
		lines [][]string
	}
	type args struct {
		fromRowIndex int
		rangeNum     int
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantNewLines [][]AAA
		wantErr      bool
	}{
		{
			name: "GetAAALineRange",
			fields: fields{
				lines: [][]string{{"1", "2", "3"}, {"4", "5", "6"}},
			},
			args: args{
				fromRowIndex: 0,
				rangeNum:     1,
			},
			wantNewLines: [][]AAA{{1, 2, 3}},
			wantErr:      false,
		},
		{
			name: "GetAAALineRange",
			fields: fields{
				lines: [][]string{{"1", "2", "3"}, {"4", "5", "6"}},
			},
			args: args{
				fromRowIndex: 0,
				rangeNum:     2,
			},
			wantNewLines: [][]AAA{{1, 2, 3}, {4, 5, 6}},
			wantErr:      false,
		},
		{
			name: "GetAAALineRange",
			fields: fields{
				lines: [][]string{{"1", "2", "3"}, {"4", "5", "6"}},
			},
			args: args{
				fromRowIndex: 1,
				rangeNum:     1,
			},
			wantNewLines: [][]AAA{{4, 5, 6}},
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Input{
				lines: tt.fields.lines,
			}
			gotNewLines, err := i.GetAAALineRange(tt.args.fromRowIndex, tt.args.rangeNum)
			if (err != nil) != tt.wantErr {
				t.Errorf("Input.GetAAALineRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotNewLines, tt.wantNewLines) {
				t.Errorf("Input.GetAAALineRange() = %v, want %v", gotNewLines, tt.wantNewLines)
			}
		})
	}
}
