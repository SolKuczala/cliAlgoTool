package customsort

import (
	"encoding/csv"
	"os"
	"testing"
)

func BenchmarkFindOptimalPath(b *testing.B) {
	file, err := os.Open("../benchmark-input/benchmark_input.csv")
	if err != nil {
		b.Fatal("failed to open file")
	}
	defer file.Close()
	reader := csv.NewReader(file)
	for i := 0; i < b.N; i++ {
		_, err := FindOptimalPath(reader, 1)
		if err != nil {
			b.Fatal("failed to find optimal path")
		}
	}
}
