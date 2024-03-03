package hezzl_test

import (
	"testing"

	"github.com/romanchechyotkin/hezzl-test-task"
)

func TestAdd(t *testing.T) {
	sum := hezzl.Add(1, 2)
	if sum != 3 {
		t.Fatal()
	}
}