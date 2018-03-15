package cfmt

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/fatih/color"
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

func TestContextualFprintf(t *testing.T) {
	rb := new(bytes.Buffer)
	color.Output = rb
	color.NoColor = false

	testCases := []struct {
		cfmt FprintfFunc
		code color.Attribute
		text string
	}{
		{cfmt: Fsuccessf, code: color.FgGreen, text: "success"},
		{cfmt: Finfof, code: color.FgCyan, text: "info"},
		{cfmt: Fwarningf, code: color.FgYellow, text: "warning"},
		{cfmt: Ferrorf, code: color.FgRed, text: "error"},
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
	color.Output = rb
	color.NoColor = false

	testCases := []struct {
		cfmt PrintfFunc
		code color.Attribute
		text string
	}{
		{cfmt: Successf, code: color.FgGreen, text: "success"},
		{cfmt: Infof, code: color.FgCyan, text: "info"},
		{cfmt: Warningf, code: color.FgYellow, text: "warning"},
		{cfmt: Errorf, code: color.FgRed, text: "error"},
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
		code color.Attribute
		text string
	}{
		{cfmt: Ssuccessf, code: color.FgGreen, text: "success"},
		{cfmt: Sinfof, code: color.FgCyan, text: "info"},
		{cfmt: Swarningf, code: color.FgYellow, text: "warning"},
		{cfmt: Serrorf, code: color.FgRed, text: "error"},
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
	color.Output = rb
	color.NoColor = false

	testCases := []struct {
		cfmt FprintFunc
		code color.Attribute
		text string
	}{
		{cfmt: Fsuccess, code: color.FgGreen, text: "success"},
		{cfmt: Finfo, code: color.FgCyan, text: "info"},
		{cfmt: Fwarning, code: color.FgYellow, text: "warning"},
		{cfmt: Ferror, code: color.FgRed, text: "error"},
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
	color.Output = rb
	color.NoColor = false

	testCases := []struct {
		cfmt PrintFunc
		code color.Attribute
		text string
	}{
		{cfmt: Success, code: color.FgGreen, text: "success"},
		{cfmt: Info, code: color.FgCyan, text: "info"},
		{cfmt: Warning, code: color.FgYellow, text: "warning"},
		{cfmt: Error, code: color.FgRed, text: "error"},
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
		code color.Attribute
		text string
	}{
		{cfmt: Ssuccess, code: color.FgGreen, text: "success"},
		{cfmt: Sinfo, code: color.FgCyan, text: "info"},
		{cfmt: Swarning, code: color.FgYellow, text: "warning"},
		{cfmt: Serror, code: color.FgRed, text: "error"},
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
	color.Output = rb
	color.NoColor = false

	testCases := []struct {
		cfmt FprintlnFunc
		code color.Attribute
		text string
	}{
		{cfmt: Fsuccessln, code: color.FgGreen, text: "success"},
		{cfmt: Finfoln, code: color.FgCyan, text: "info"},
		{cfmt: Fwarningln, code: color.FgYellow, text: "warning"},
		{cfmt: Ferrorln, code: color.FgRed, text: "error"},
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
	color.Output = rb
	color.NoColor = false

	testCases := []struct {
		cfmt PrintlnFunc
		code color.Attribute
		text string
	}{
		{cfmt: Successln, code: color.FgGreen, text: "success"},
		{cfmt: Infoln, code: color.FgCyan, text: "info"},
		{cfmt: Warningln, code: color.FgYellow, text: "warning"},
		{cfmt: Errorln, code: color.FgRed, text: "error"},
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
		code color.Attribute
		text string
	}{
		{cfmt: Ssuccessln, code: color.FgGreen, text: "success"},
		{cfmt: Sinfoln, code: color.FgCyan, text: "info"},
		{cfmt: Swarningln, code: color.FgYellow, text: "warning"},
		{cfmt: Serrorln, code: color.FgRed, text: "error"},
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
