package mono

import (
	"testing"
)

func TestNewMono_solver(t *testing.T) {
	ms := NewMono_Solver()
	if ms.max_top_ != 1 {
		t.Error("ms.max_top_ should be 1")
	}
}
