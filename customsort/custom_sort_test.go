package customsort

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	type args struct {
		column int
		matrix [][]string
	}

	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "test case 1",
			args: args{
				column: 2,
				matrix: [][]string{
					{"", "", "AA 1"},
					{"", "", "A 1"},
				},
			},
			want: [][]string{
				{"", "", "A 1"},
				{"", "", "AA 1"},
			},
		},
		{
			name: "test case 2",
			args: args{
				column: 2,
				matrix: [][]string{
					{"", "", "AA 1"},
					{"", "", "Z 1"},
				},
			},
			want: [][]string{
				{"", "", "Z 1"},
				{"", "", "AA 1"},
			},
		},
		{
			name: "test case 3",
			args: args{
				column: 2,
				matrix: [][]string{
					{"", "", "Z 2"},
					{"", "", "Z 1"},
				},
			},
			want: [][]string{
				{"", "", "Z 1"},
				{"", "", "Z 2"},
			},
		},
		{
			name: "test case 4",
			args: args{
				column: 2,
				matrix: [][]string{
					{"", "", "A 10"},
					{"", "", "A 1"},
				},
			},
			want: [][]string{
				{"", "", "A 1"},
				{"", "", "A 10"},
			},
		},
		{
			name: "test case 4",
			args: args{
				column: 2,
				matrix: [][]string{
					{"", "", "AA 10"},
					{"", "", "AA 1"},
				},
			},
			want: [][]string{
				{"", "", "AA 1"},
				{"", "", "AA 10"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SortCSVbyColumnIdx(tt.args.column, tt.args.matrix)

			if reflect.DeepEqual(tt.want, tt.args.matrix) == false {
				t.Errorf("got %v, want %v", tt.args.matrix, tt.want)
			}
		})
	}
}
