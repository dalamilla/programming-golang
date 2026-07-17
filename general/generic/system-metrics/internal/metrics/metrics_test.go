package metrics

import (
	"runtime"
	"testing"
)

func TestGetSystemMetrics(t *testing.T) {
	expected := "linux"

	got, _ := GetSystemMetrics()
	if got.OS != expected {
		t.Errorf("GetSystemMetrics().OS => %v, should return %v", got.OS, expected)
	}
}

func TestGetCPU(t *testing.T) {
	expected := runtime.NumCPU()

	got, _ := GetCPU()
	if got.Cores != expected {
		t.Errorf("GetCPU().Cores => %v, should return %v", got.Cores, expected)
	}

}
