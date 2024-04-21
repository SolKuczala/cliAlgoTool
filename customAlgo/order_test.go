package customAlgo

import "testing"

func TestOrder(t *testing.T) {
	type args struct {
		column int
		matrix [][]string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			OrderCSVbyColumnIdx(tt.args.column, tt.args.matrix)
		})
	}
}
