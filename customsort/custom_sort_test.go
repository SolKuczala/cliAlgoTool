package customsort

import (
	"bytes"
	"clialgotool/utils"
	"encoding/csv"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindOptimalPath(t *testing.T) {
	// create orderawaremaps want cases
	want1 := &utils.OrderAwareMap{
		CSVdata: map[string]utils.CSVRow{
			" A01012345": {ProductCode: "12345", Quantity: 1, PickLocation: "A 1"},
			" A01023456": {ProductCode: "23456", Quantity: 1, PickLocation: "A 1"},
		},
	}
	want1.SetOrder([]string{" A01012345", " A01023456"})

	want3 := &utils.OrderAwareMap{
		CSVdata: map[string]utils.CSVRow{
			" A01012345": {ProductCode: "12345", Quantity: 1, PickLocation: "A 1"},
			" A02023456": {ProductCode: "23456", Quantity: 1, PickLocation: "A 2"},
		},
	}
	want3.SetOrder([]string{" A01012345", " A02023456"})

	want4 := &utils.OrderAwareMap{
		CSVdata: map[string]utils.CSVRow{
			" A01012345": {ProductCode: "12345", Quantity: 1, PickLocation: "A 1"},
			" A01023456": {ProductCode: "23456", Quantity: 1, PickLocation: "A 1"},
		},
	}
	want4.SetOrder([]string{" A01012345", " A01023456"})

	want5 := &utils.OrderAwareMap{
		CSVdata: map[string]utils.CSVRow{
			" A01012345": {ProductCode: "12345", Quantity: 1, PickLocation: "A 1"},
			"AA01056789": {ProductCode: "56789", Quantity: 1, PickLocation: "AA 1"},
			" C09034567": {ProductCode: "34567", Quantity: 1, PickLocation: "C 9"},
			"AA05023456": {ProductCode: "23456", Quantity: 1, PickLocation: "AA 5"},
			"AZ02056789": {ProductCode: "56789", Quantity: 1, PickLocation: "AZ 2"},
			"AZ10045678": {ProductCode: "45678", Quantity: 1, PickLocation: "AZ 10"},
		},
	}
	want5.SetOrder([]string{" A01012345", " C09034567", "AA01056789", "AA05023456", "AZ02056789", "AZ10045678"})

	// push oamts to array to set headers in at once
	oamts := []*utils.OrderAwareMap{want1, want3, want4, want5}
	for _, oamt := range oamts {
		oamt.SetHeader([]string{"product_code", "quantity", "pick_location"})
	}

	type args struct {
		input io.Reader
		skip  int
	}
	tests := []struct {
		name    string
		args    args
		want    *utils.OrderAwareMap
		wantErr bool
	}{
		{
			name: "1 - same location, different product code",
			args: args{
				input: bytes.NewReader([]byte(
					"product_code,quantity,pick_location\n12345,1,A 1\n23456,1,A 1\n")),
				skip: 1,
			},
			want:    want1,
			wantErr: false,
		},
		{
			name: "3 - different product code, different location",
			args: args{
				input: bytes.NewReader([]byte(
					"product_code,quantity,pick_location\n12345,1,A 1\n23456,1,A 2\n")),
				skip: 1,
			},
			want:    want3,
			wantErr: false,
		},
		{
			name: "4 - different product code, same location",
			args: args{
				input: bytes.NewReader([]byte(
					"product_code,quantity,pick_location\n12345,1,A 1\n23456,1,A 1\n")),
				skip: 1,
			},
			want:    want4,
			wantErr: false,
		},
		{
			name: "5 - correct order by alphabet and number(edge cases)",
			args: args{
				input: bytes.NewReader([]byte(
					"product_code,quantity,pick_location\n12345,1,A 1\n23456,1,AA 5\n34567,1,C 9\n45678,1,AZ 10\n56789,1,AA 1\n56789,1,AZ 2\n")),
				skip: 1,
			},
			want:    want5,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := csv.NewReader(tt.args.input)
			got, err := FindOptimalPath(reader, tt.args.skip)
			if (err != nil) != tt.wantErr {
				assert.EqualError(t, err, "failed to read input:EOF")
			}
			assert.Equal(t, tt.want.GetOrder(), got.GetOrder())
			assert.Equal(t, tt.want.GetHeader(), got.GetHeader())
			assert.Equal(t, tt.want.CSVdata, got.CSVdata)
		})
	}
}
