package cfmt

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

type (
	FprintfFunc  func(io.Writer, string, ...interface{}) (int, error)
	PrintfFunc   func(string, ...interface{}) (int, error)
	SprintfFunc  func(string, ...interface{}) string
	FprintFunc   func(io.Writer, ...interface{}) (int, error)
	PrintFunc    func(...interface{}) (int, error)
	SprintFunc   func(...interface{}) string
	FprintlnFunc func(io.Writer, ...interface{}) (int, error)
	PrintlnFunc  func(...interface{}) (int, error)
	SprintlnFunc func(...interface{}) string
)

const (
	Green  = 32
	Cyan   = 36
	Yellow = 33
	Red    = 31
)

func TestContextualFprintf(t *testing.T) {
	rb := new(bytes.Buffer)

	testCases := []struct {
		cfmt FprintfFunc
		code int
		text string
	}{
		{cfmt: Fsuccessf, code: Green, text: "success"},
		{cfmt: Finfof, code: Cyan, text: "info"},
		{cfmt: Fwarningf, code: Yellow, text: "warning"},
		{cfmt: Ferrorf, code: Red, text: "error"},
	}

	for _, c := range testCases {
		c.cfmt(rb, "This is %s", c.text)
		line, _ := rb.ReadString('\n')
		scannedLine := fmt.Sprintf("%q", line)
		colored := fmt.Sprintf("\x1b[%dmThis is %s\x1b[0m", c.code, c.text)
		escapedForm := fmt.Sprintf("%q", colored)
		if scannedLine != escapedForm {
			t.Errorf("Expecting %s, got '%s'\n", escapedForm, scannedLine)
		}
	}
}

func TestContextualPrintf(t *testing.T) {
	rb := new(bytes.Buffer)
	output = rb

	testCases := []struct {
		cfmt PrintfFunc
		code int
		text string
	}{
		{cfmt: Successf, code: Green, text: "success"},
		{cfmt: Infof, code: Cyan, text: "info"},
		{cfmt: Warningf, code: Yellow, text: "warning"},
		{cfmt: Errorf, code: Red, text: "error"},
	}

	for _, c := range testCases {
		c.cfmt("This is %s", c.text)
		line, _ := rb.ReadString('\n')
		scannedLine := fmt.Sprintf("%q", line)
		colored := fmt.Sprintf("\x1b[%dmThis is %s\x1b[0m", c.code, c.text)
		escapedForm := fmt.Sprintf("%q", colored)
		if scannedLine != escapedForm {
			t.Errorf("Expecting %s, got '%s'\n", escapedForm, scannedLine)
		}
	}
}

func TestContextualSprintf(t *testing.T) {
	testCases := []struct {
		cfmt SprintfFunc
		code int
		text string
	}{
		{cfmt: Ssuccessf, code: Green, text: "success"},
		{cfmt: Sinfof, code: Cyan, text: "info"},
		{cfmt: Swarningf, code: Yellow, text: "warning"},
		{cfmt: Serrorf, code: Red, text: "error"},
	}

	for _, c := range testCases {
		line := c.cfmt("This is %s", c.text)
		scannedLine := fmt.Sprintf("%q", line)
		colored := fmt.Sprintf("\x1b[%dmThis is %s\x1b[0m", c.code, c.text)
		escapedForm := fmt.Sprintf("%q", colored)
		if scannedLine != escapedForm {
			t.Errorf("Expecting %s, got '%s'\n", escapedForm, scannedLine)
		}
	}
}

func TestContextualFprint(t *testing.T) {
	rb := new(bytes.Buffer)
	output = rb

	testCases := []struct {
		cfmt FprintFunc
		code int
		text string
	}{
		{cfmt: Fsuccess, code: Green, text: "success"},
		{cfmt: Finfo, code: Cyan, text: "info"},
		{cfmt: Fwarning, code: Yellow, text: "warning"},
		{cfmt: Ferror, code: Red, text: "error"},
	}

	for _, c := range testCases {
		c.cfmt(rb, c.text)
		line, _ := rb.ReadString('\n')
		scannedLine := fmt.Sprintf("%q", line)
		colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", c.code, c.text)
		escapedForm := fmt.Sprintf("%q", colored)
		if scannedLine != escapedForm {
			t.Errorf("Expecting %s, got '%s'\n", escapedForm, scannedLine)
		}
	}
}

func TestContextualPrint(t *testing.T) {
	rb := new(bytes.Buffer)
	output = rb

	testCases := []struct {
		cfmt PrintFunc
		code int
		text string
	}{
		{cfmt: Success, code: Green, text: "success"},
		{cfmt: Info, code: Cyan, text: "info"},
		{cfmt: Warning, code: Yellow, text: "warning"},
		{cfmt: Error, code: Red, text: "error"},
	}

	for _, c := range testCases {
		c.cfmt(c.text)
		line, _ := rb.ReadString('\n')
		scannedLine := fmt.Sprintf("%q", line)
		colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", c.code, c.text)
		escapedForm := fmt.Sprintf("%q", colored)
		if scannedLine != escapedForm {
			t.Errorf("Expecting %s, got '%s'\n", escapedForm, scannedLine)
		}
	}
}

func TestContextualSprint(t *testing.T) {
	testCases := []struct {
		cfmt SprintFunc
		code int
		text string
	}{
		{cfmt: Ssuccess, code: Green, text: "success"},
		{cfmt: Sinfo, code: Cyan, text: "info"},
		{cfmt: Swarning, code: Yellow, text: "warning"},
		{cfmt: Serror, code: Red, text: "error"},
	}

	for _, c := range testCases {
		line := c.cfmt(c.text)
		scannedLine := fmt.Sprintf("%q", line)
		colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", c.code, c.text)
		escapedForm := fmt.Sprintf("%q", colored)
		if scannedLine != escapedForm {
			t.Errorf("Expecting %s, got '%s'\n", escapedForm, scannedLine)
		}
	}
}

func TestContextualFprintln(t *testing.T) {
	rb := new(bytes.Buffer)
	output = rb

	testCases := []struct {
		cfmt FprintlnFunc
		code int
		text string
	}{
		{cfmt: Fsuccessln, code: Green, text: "success"},
		{cfmt: Finfoln, code: Cyan, text: "info"},
		{cfmt: Fwarningln, code: Yellow, text: "warning"},
		{cfmt: Ferrorln, code: Red, text: "error"},
	}

	for _, c := range testCases {
		c.cfmt(rb, c.text)
		line, _ := rb.ReadString(' ')
		scannedLine := fmt.Sprintf("%q", line)
		colored := fmt.Sprintf("\x1b[%dm%s\n\x1b[0m", c.code, c.text)
		escapedForm := fmt.Sprintf("%q", colored)
		if scannedLine != escapedForm {
			t.Errorf("Expecting %s, got '%s'\n", escapedForm, scannedLine)
		}
	}
}

func TestContextualPrintln(t *testing.T) {
	rb := new(bytes.Buffer)
	output = rb

	testCases := []struct {
		cfmt PrintlnFunc
		code int
		text string
	}{
		{cfmt: Successln, code: Green, text: "success"},
		{cfmt: Infoln, code: Cyan, text: "info"},
		{cfmt: Warningln, code: Yellow, text: "warning"},
		{cfmt: Errorln, code: Red, text: "error"},
	}

	for _, c := range testCases {
		c.cfmt(c.text)
		line, _ := rb.ReadString(' ')
		scannedLine := fmt.Sprintf("%q", line)
		colored := fmt.Sprintf("\x1b[%dm%s\n\x1b[0m", c.code, c.text)
		escapedForm := fmt.Sprintf("%q", colored)
		if scannedLine != escapedForm {
			t.Errorf("Expecting %s, got '%s'\n", escapedForm, scannedLine)
		}
	}
}

func TestContextualSprintln(t *testing.T) {
	testCases := []struct {
		cfmt SprintlnFunc
		code int
		text string
	}{
		{cfmt: Ssuccessln, code: Green, text: "success"},
		{cfmt: Sinfoln, code: Cyan, text: "info"},
		{cfmt: Swarningln, code: Yellow, text: "warning"},
		{cfmt: Serrorln, code: Red, text: "error"},
	}

	for _, c := range testCases {
		line := c.cfmt(c.text)
		scannedLine := fmt.Sprintf("%q", line)
		colored := fmt.Sprintf("\x1b[%dm%s\n\x1b[0m", c.code, c.text)
		escapedForm := fmt.Sprintf("%q", colored)
		if scannedLine != escapedForm {
			t.Errorf("Expecting %s, got '%s'\n", escapedForm, scannedLine)
		}
	}
}
