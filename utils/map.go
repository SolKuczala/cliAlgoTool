package utils

import (
	"fmt"
	"sort"
)

type OrderAwareMap struct {
	order   []string
	CSVdata map[string]CSVRow
	header  []string
}

func (m *OrderAwareMap) SortKeys() {
	// creating ordered keys slice
	sortedKeys := make([]string, 0, len(m.CSVdata))
	for key := range m.CSVdata {
		sortedKeys = append(sortedKeys, key)
	}
	sort.Strings(sortedKeys)
	m.order = sortedKeys
}

func (m *OrderAwareMap) GetOrder() []string {
	return m.order
}

func (m *OrderAwareMap) GetHeader() []string {
	return m.header
}

func (m *OrderAwareMap) SetHeader(header []string) {
	m.header = header
}

func (m OrderAwareMap) String() string {
	return fmt.Sprintf("order:%v, values:%v", m.order, m.CSVdata)
}
