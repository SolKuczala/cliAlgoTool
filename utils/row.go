package utils

import "fmt"

type CSVRow struct {
	ProductCode  string
	Quantity     int
	PickLocation string
}

func (r CSVRow) String() string {
	return fmt.Sprintf("%v,%v,%v", r.ProductCode, r.Quantity, r.PickLocation)
}
