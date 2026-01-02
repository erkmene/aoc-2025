package utils

import (
	"testing"
)

func TestImportMultiDimensionalData(t *testing.T) {
	result := ImportMultiDimensionalData("utils_test_multidimensional_input.dat", []string{",", "-"})

	// Root node should have "," as separator
	if result.Separator != "," {
		t.Errorf("expected root separator to be %q, got %q", ",", result.Separator)
	}

	// Should have 11 comma-separated children
	if len(result.Children) != 11 {
		t.Errorf("expected 11 children, got %d", len(result.Children))
	}

	// First child should be "11-22" with "-" as separator
	firstChild := result.Children[0]
	if firstChild.Value != "11-22" {
		t.Errorf("expected first child value to be %q, got %q", "11-22", firstChild.Value)
	}
	if firstChild.Separator != "-" {
		t.Errorf("expected first child separator to be %q, got %q", "-", firstChild.Separator)
	}

	// First child should have 2 dash-separated children: "11" and "22"
	if len(firstChild.Children) != 2 {
		t.Errorf("expected first child to have 2 children, got %d", len(firstChild.Children))
	}
	if firstChild.Children[0].Value != "11" {
		t.Errorf("expected first grandchild value to be %q, got %q", "11", firstChild.Children[0].Value)
	}
	if firstChild.Children[1].Value != "22" {
		t.Errorf("expected second grandchild value to be %q, got %q", "22", firstChild.Children[1].Value)
	}

	// Last child should be "2121212118-2121212124"
	lastChild := result.Children[10]
	if lastChild.Value != "2121212118-2121212124" {
		t.Errorf("expected last child value to be %q, got %q", "2121212118-2121212124", lastChild.Value)
	}
	if lastChild.Children[0].Value != "2121212118" {
		t.Errorf("expected last grandchild first value to be %q, got %q", "2121212118", lastChild.Children[0].Value)
	}
	if lastChild.Children[1].Value != "2121212124" {
		t.Errorf("expected last grandchild second value to be %q, got %q", "2121212124", lastChild.Children[1].Value)
	}
}
