package part2

import (
	"bytes"
	"testing"
)

func TestPart1(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`

	result, err := Run(bytes.NewBufferString(input))
	if err != nil {
		t.Fatalf("Error creating buffer: %v", err)
	}

	expected := 11
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
