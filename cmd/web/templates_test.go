package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHumanDate(t *testing.T) {
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2023, 1, 18, 13, 37, 0, 0, time.UTC),
			want: "Wed 18 Jan 2023 at 13:37",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2023, 1, 18, 13, 37, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "Wed 18 Jan 2023 at 12:37",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := humanDate(tt.tm)
			assert.Equal(t, tt.want, got)
		})
	}
}
