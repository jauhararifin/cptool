package executioner

import (
	"context"
	"testing"
	"time"
)

func TestOsCommandRun(t *testing.T) {
	exec := NewOSExec()
	cmd := exec.Command("sleep", "1")
	start := time.Now()
	err := cmd.Run()
	duration := time.Since(start)
	if err != nil {
		t.Error(err)
	}
	if duration.Seconds() < 1 {
		t.Error("Run should wait 1 second")
	}
}

func TestOsCommandRunWithContext(t *testing.T) {
	exec := NewOSExec()
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	cmd := exec.CommandContext(ctx, "sleep", "1")
	start := time.Now()
	err := cmd.Run()
	duration := time.Since(start)
	if err == nil {
		t.Error("Run should returns error")
	}
	if duration.Seconds() > 700 {
		t.Error("Run should wait no more than 500 ms")
	}
}
