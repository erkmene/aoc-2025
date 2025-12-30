package utils

import "testing"

func TestImportNewlineSeparatedData(t *testing.T) {
	result := ImportNewlineSeparatedData("../01.test.dat")

	if len(result) != 10 {
		t.Errorf("expected 10 lines, got %d", len(result))
	}

	expectedFirst := "L68"
	if result[0] != expectedFirst {
		t.Errorf("expected first line to be %q, got %q", expectedFirst, result[0])
	}

	expectedLast := "L82"
	if result[9] != expectedLast {
		t.Errorf("expected last line to be %q, got %q", expectedLast, result[9])
	}
}
