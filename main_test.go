package main

import (
	"testing"
	"time"
)

func TestDaysLeft(t *testing.T) {
	target := time.Date(2027, time.January, 1, 0, 0, 0, 0, time.Local)
	tests := []struct {
		name     string
		mockNow  time.Time
		expected int
	}{
		{
			name:     "Ровно за неделю до Нового Года",
			mockNow:  time.Date(2026, time.December, 25, 0, 0, 0, 0, time.Local),
			expected: 7,
		},
		{
			name:     "Осталось меньше суток (12 часов) -> округление в меньшую сторону",
			mockNow:  time.Date(2026, time.December, 31, 12, 0, 0, 0, time.Local),
			expected: 0,
		},
		{
			name:     "Новый год уже наступил (прошло 2 дня)",
			mockNow:  time.Date(2027, time.January, 3, 0, 0, 0, 0, time.Local),
			expected: -2,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := DaysLeft(tc.mockNow, target)
			if got != tc.expected {
				t.Errorf("DaysLeft() во время '%s' = %d; хотим %d", tc.name, got, tc.expected)
			}
		})
	}
}
