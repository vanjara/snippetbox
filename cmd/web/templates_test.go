package main

import (
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2022, 2, 21, 10, 0, 0, 0, time.UTC),
			want: "21 Feb 2022 at 10:00",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2022, 2, 21, 10, 0, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "21 Feb 2022 at 09:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := humanDate(tt.tm)
			if got != tt.want {
				t.Errorf("want %q; got %q", tt.want, got)
			}
		})
	}
}
