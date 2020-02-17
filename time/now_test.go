package time

import (
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "first"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Now()
			t.Log(time.Since(time.Now().Add(2*time.Minute)) < 0)
		})
	}
}
