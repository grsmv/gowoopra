package gowoopra

import (
	//"reflect"
	"testing"
)

func TestNewTracker(t *testing.T) {
	_, err := NewTracker(map[string]string{
		"host": "",
	})

	if err == nil {
		t.Error("Error expected during NewTracker creation without a host")
	}
}
