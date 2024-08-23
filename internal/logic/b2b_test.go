package logic

import (
	"testing"

	"github.com/mrRestyling/ExchangeRest/internal/models"
)

func TestBusiness(t *testing.T) {
	tests := []struct {
		name     string
		request  *models.Request
		expected [][]int
	}{
		{
			name: "simple case",
			request: &models.Request{
				Amount:    10,
				Banknotes: []int{5, 5},
			},
			expected: [][]int{{5, 5}},
		},
		{
			name: "multiple combinations",
			request: &models.Request{
				Amount:    10,
				Banknotes: []int{2, 2, 2, 2, 2},
			},
			expected: [][]int{{2, 2, 2, 2, 2}},
		},
		{
			name: "no combinations",
			request: &models.Request{
				Amount:    10,
				Banknotes: []int{3, 3, 3},
			},
			expected: [][]int{},
		},
		{
			name: "empty banknotes",
			request: &models.Request{
				Amount:    10,
				Banknotes: []int{},
			},
			expected: [][]int{},
		},
		{
			name: "zero amount",
			request: &models.Request{
				Amount:    0,
				Banknotes: []int{5, 5},
			},
			expected: [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := (&Appeal{}).Business(tt.request)
			if !equal(actual, tt.expected) {
				t.Errorf("Business() = %v, want %v", actual, tt.expected)
			}
		})
	}
}

func equal(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if !equalSlices(v, b[i]) {
			return false
		}
	}
	return true
}

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
