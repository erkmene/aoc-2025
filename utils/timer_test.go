package utils

import (
	"bytes"
	"os"
	"strings"
	"testing"
	"time"
)

func TestNewTimeRecord(t *testing.T) {
	before := time.Now()
	tr := NewTimeRecord()
	after := time.Now()

	// Start time should be between before and after
	if tr.Start.Before(before) || tr.Start.After(after) {
		t.Errorf("Start time %v should be between %v and %v", tr.Start, before, after)
	}

	// Laps should be empty
	if len(tr.Laps) != 0 {
		t.Errorf("expected 0 laps, got %d", len(tr.Laps))
	}
}

func TestLap(t *testing.T) {
	tr := NewTimeRecord()

	// Wait a bit then record first lap
	time.Sleep(10 * time.Millisecond)
	totalDuration1, lapDuration1 := tr.Lap()

	// First lap: total and lap duration should be approximately equal
	if totalDuration1 < 10*time.Millisecond {
		t.Errorf("first lap total duration %v should be >= 10ms", totalDuration1)
	}
	if lapDuration1 < 10*time.Millisecond {
		t.Errorf("first lap duration %v should be >= 10ms", lapDuration1)
	}

	// Should have 1 lap recorded
	if len(tr.Laps) != 1 {
		t.Errorf("expected 1 lap, got %d", len(tr.Laps))
	}

	// Wait a bit then record second lap
	time.Sleep(20 * time.Millisecond)
	totalDuration2, lapDuration2 := tr.Lap()

	// Second lap: total should be >= 30ms, lap should be >= 20ms
	if totalDuration2 < 30*time.Millisecond {
		t.Errorf("second lap total duration %v should be >= 30ms", totalDuration2)
	}
	if lapDuration2 < 20*time.Millisecond {
		t.Errorf("second lap duration %v should be >= 20ms", lapDuration2)
	}

	// Total should be greater than first total
	if totalDuration2 <= totalDuration1 {
		t.Errorf("second total %v should be greater than first total %v", totalDuration2, totalDuration1)
	}

	// Should have 2 laps recorded
	if len(tr.Laps) != 2 {
		t.Errorf("expected 2 laps, got %d", len(tr.Laps))
	}
}

func TestString(t *testing.T) {
	tr := NewTimeRecord()
	str := tr.String()

	// Should contain "Start:" and the RFC3339 formatted time
	if !strings.Contains(str, "Start:") {
		t.Errorf("String() should contain 'Start:', got %q", str)
	}
	if !strings.Contains(str, "Laps:") {
		t.Errorf("String() should contain 'Laps:', got %q", str)
	}
}

func TestLapAndLog(t *testing.T) {
	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	tr := NewTimeRecord()
	time.Sleep(5 * time.Millisecond)
	tr.LapAndLog("Test message")

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	// Should contain the message
	if !strings.Contains(output, "Test message:") {
		t.Errorf("LapAndLog output should contain 'Test message:', got %q", output)
	}
	// Should contain "last lap" and "total"
	if !strings.Contains(output, "last lap") {
		t.Errorf("LapAndLog output should contain 'last lap', got %q", output)
	}
	if !strings.Contains(output, "total") {
		t.Errorf("LapAndLog output should contain 'total', got %q", output)
	}

	// Should have recorded a lap
	if len(tr.Laps) != 1 {
		t.Errorf("expected 1 lap after LapAndLog, got %d", len(tr.Laps))
	}
}
