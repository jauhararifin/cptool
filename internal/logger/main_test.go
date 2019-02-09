package logger

import (
	"bytes"
	"strings"
	"testing"
)

func TestPrintInfo(t *testing.T) {
	buf := new(bytes.Buffer)
	logger := New(buf, INFO)
	logger.PrintInfo("just", "some", "test")
	str := buf.String()
	if !strings.Contains(str, "Info") || !strings.Contains(str, "justsometest") {
		t.Fail()
	}
}

func TestPrintInfoWithHigherLoggingLevel(t *testing.T) {
	buf := new(bytes.Buffer)
	logger := New(buf, INFO+1)
	logger.PrintInfo("just", "some", "test")
	str := buf.String()
	if len(str) > 0 {
		t.Fail()
	}
}

func TestPrintWarning(t *testing.T) {
	buf := new(bytes.Buffer)
	logger := New(buf, WARN)
	logger.PrintWarning("just", "some", "test")
	str := buf.String()
	if !strings.Contains(str, "Warn") || !strings.Contains(str, "justsometest") {
		t.Fail()
	}
}

func TestPrintWarningWithHigherLoggingLevel(t *testing.T) {
	buf := new(bytes.Buffer)
	logger := New(buf, WARN+1)
	logger.PrintWarning("just", "some", "test")
	str := buf.String()
	if len(str) > 0 {
		t.Fail()
	}
}

func TestPrintError(t *testing.T) {
	buf := new(bytes.Buffer)
	logger := New(buf, ERROR)
	logger.PrintError("just", "some", "test")
	str := buf.String()
	if !strings.Contains(str, "Error") || !strings.Contains(str, "justsometest") {
		t.Fail()
	}
}

func TestPrintErrorWithHigherLoggingLevel(t *testing.T) {
	buf := new(bytes.Buffer)
	logger := New(buf, ERROR+1)
	logger.PrintError("just", "some", "test")
	str := buf.String()
	if len(str) > 0 {
		t.Fail()
	}
}

func TestPrintSuccess(t *testing.T) {
	buf := new(bytes.Buffer)
	logger := New(buf, INFO)
	logger.PrintSuccess("just", "some", "test")
	str := buf.String()
	if !strings.Contains(str, "Success") || !strings.Contains(str, "justsometest") {
		t.Fail()
	}
}

func TestPrintSuccessWithHigherLoggingLevel(t *testing.T) {
	buf := new(bytes.Buffer)
	logger := New(buf, INFO+1)
	logger.PrintSuccess("just", "some", "test")
	str := buf.String()
	if len(str) > 0 {
		t.Fail()
	}
}

func TestPrint(t *testing.T) {
	buf := new(bytes.Buffer)
	logger := New(buf, 10)
	logger.Print(11, "just", "some", "test")
	str := buf.String()
	if !strings.Contains(str, "justsometest") {
		t.Fail()
	}
}

func TestPrintWithHigherLevel(t *testing.T) {
	buf := new(bytes.Buffer)
	logger := New(buf, 10)
	logger.Print(9, "just", "some", "test")
	str := buf.String()
	if len(str) > 0 {
		t.Fail()
	}
}

func TestPrintln(t *testing.T) {
	buf := new(bytes.Buffer)
	logger := New(buf, 10)
	logger.Println(11, "just", "some", "test")
	str := buf.String()
	if !strings.Contains(str, "justsometest\n") {
		t.Fail()
	}
}

func TestPrintlnWithHigherLevel(t *testing.T) {
	buf := new(bytes.Buffer)
	logger := New(buf, 10)
	logger.Println(9, "just", "some", "test")
	str := buf.String()
	if len(str) > 0 {
		t.Fail()
	}
}

func TestPrintf(t *testing.T) {
	buf := new(bytes.Buffer)
	logger := New(buf, 10)
	logger.Printf(11, "%s-%s-%s", "just", "some", "test")
	str := buf.String()
	if !strings.Contains(str, "just-some-test") {
		t.Fail()
	}
}

func TestPrintfWithHigherLevel(t *testing.T) {
	buf := new(bytes.Buffer)
	logger := New(buf, 10)
	logger.Printf(9, "%s-%s-%s", "just", "some", "test")
	str := buf.String()
	if len(str) > 0 {
		t.Fail()
	}
}
