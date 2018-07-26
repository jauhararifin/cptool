package logger

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestPrintInfo(t *testing.T) {
	buf := new(bytes.Buffer)
	OutputTarget = buf
	PrintInfo("just", "some", "test")
	str := buf.String()
	fmt.Println(str)
	if !strings.Contains(str, "Info") || !strings.Contains(str, "justsometest") {
		t.Fail()
	}
}

func TestPrintWarning(t *testing.T) {
	buf := new(bytes.Buffer)
	OutputTarget = buf
	PrintWarning("just", "some", "test")
	str := buf.String()
	if !strings.Contains(str, "Warn") || !strings.Contains(str, "justsometest") {
		t.Fail()
	}
}

func TestPrintError(t *testing.T) {
	buf := new(bytes.Buffer)
	OutputTarget = buf
	PrintError("just", "some", "test")
	str := buf.String()
	if !strings.Contains(str, "Error") || !strings.Contains(str, "justsometest") {
		t.Fail()
	}
}

func TestPrintSuccess(t *testing.T) {
	buf := new(bytes.Buffer)
	OutputTarget = buf
	PrintSuccess("just", "some", "test")
	str := buf.String()
	if !strings.Contains(str, "Success") || !strings.Contains(str, "justsometest") {
		t.Fail()
	}
}
