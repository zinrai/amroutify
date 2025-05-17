package routing

import (
	"testing"
)

func TestCompareReceivers(t *testing.T) {
	tests := []struct {
		name     string
		expected []string
		actual   []string
		want     bool
	}{
		{
			name:     "identical arrays",
			expected: []string{"receiver1", "receiver2", "receiver3"},
			actual:   []string{"receiver1", "receiver2", "receiver3"},
			want:     true,
		},
		{
			name:     "same content different order",
			expected: []string{"receiver1", "receiver2", "receiver3"},
			actual:   []string{"receiver3", "receiver1", "receiver2"},
			want:     true,
		},
		{
			name:     "different lengths",
			expected: []string{"receiver1", "receiver2", "receiver3"},
			actual:   []string{"receiver1", "receiver2"},
			want:     false,
		},
		{
			name:     "different content",
			expected: []string{"receiver1", "receiver2", "receiver3"},
			actual:   []string{"receiver1", "receiver2", "receiver4"},
			want:     false,
		},
		{
			name:     "empty arrays",
			expected: []string{},
			actual:   []string{},
			want:     true,
		},
		{
			name:     "one empty array",
			expected: []string{"receiver1"},
			actual:   []string{},
			want:     false,
		},
		{
			name:     "with duplicates in expected",
			expected: []string{"receiver1", "receiver1", "receiver2"},
			actual:   []string{"receiver1", "receiver2", "receiver1"},
			want:     true,
		},
		{
			name:     "with duplicates but different counts",
			expected: []string{"receiver1", "receiver1", "receiver2"},
			actual:   []string{"receiver1", "receiver2", "receiver2"},
			want:     false,
		},
		{
			name:     "case sensitivity check",
			expected: []string{"Receiver1", "receiver2"},
			actual:   []string{"receiver1", "Receiver2"},
			want:     false,
		},
		{
			name:     "subset check",
			expected: []string{"receiver1", "receiver2"},
			actual:   []string{"receiver1", "receiver2", "receiver3"},
			want:     false,
		},
		{
			name:     "single element arrays",
			expected: []string{"receiver1"},
			actual:   []string{"receiver1"},
			want:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareReceivers(tt.expected, tt.actual); got != tt.want {
				t.Errorf("CompareReceivers() = %v, want %v\nExpected: %v\nActual: %v",
					got, tt.want, tt.expected, tt.actual)
			}
		})
	}
}
