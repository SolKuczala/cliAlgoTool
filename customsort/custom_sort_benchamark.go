package customsort

import (
	"bytes"
	"encoding/csv"
	"testing"
)

// create benchmarks for the main function

func BenchmarkFindOptimalPath(b *testing.B) {
	// create input for the function
	input := bytes.NewReader([]byte(
		"product_code,quantity,pick_location\n12345,1,A 1\n23456,1,AA 5\n34567,1,C 9\n45678,1,AZ 10\n56789,1,AA 1\n56789,1,AZ 2\n"))
	reader := csv.NewReader(input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindOptimalPath(reader, 1)
	}
}
