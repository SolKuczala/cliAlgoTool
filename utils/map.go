package utils

import (
	"fmt"
	"sort"
)

type OrderAwareMap struct {
	order  []string
	Values map[string]CSVRow
}

func (m *OrderAwareMap) SortKeys() {
	sortedKeys := make([]string, 0, len(m.Values))
	for key := range m.Values {
		sortedKeys = append(sortedKeys, key)
	}
	sort.Strings(sortedKeys)
	m.order = sortedKeys
}

func (m *OrderAwareMap) GetOrder() []string {
	return m.order
}

func (m OrderAwareMap) String() string {
	return fmt.Sprintf("order:%v, values:%v", m.order, m.Values)
}
