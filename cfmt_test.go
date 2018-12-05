package cfmt

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/logrusorgru/aurora"
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

	testCases := []struct {
		cfmt  FprintfFunc
		color aurora.Color
		text  string
	}{
		{cfmt: Fsuccessf, color: aurora.Color(aurora.GreenFg), text: "success"},
		{cfmt: Finfof, color: aurora.Color(aurora.CyanFg), text: "info"},
		{cfmt: Fwarningf, color: aurora.Color(aurora.BrownFg), text: "warning"},
		{cfmt: Ferrorf, color: aurora.Color(aurora.RedFg), text: "error"},
	}

	for _, c := range testCases {
		c.cfmt(rb, "This is %s", c.text)
		line, _ := rb.ReadString('\n')
		scannedLine := fmt.Sprintf("%q", line)
		colored := fmt.Sprintf("\x1b[%smThis is %s\x1b[0m", c.color.Nos(), c.text)
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
		cfmt  PrintfFunc
		color aurora.Color
		text  string
	}{
		{cfmt: Successf, color: aurora.Color(aurora.GreenFg), text: "success"},
		{cfmt: Infof, color: aurora.Color(aurora.CyanFg), text: "info"},
		{cfmt: Warningf, color: aurora.Color(aurora.BrownFg), text: "warning"},
		{cfmt: Errorf, color: aurora.Color(aurora.RedFg), text: "error"},
	}

	for _, c := range testCases {
		c.cfmt("This is %s", c.text)
		line, _ := rb.ReadString('\n')
		scannedLine := fmt.Sprintf("%q", line)
		colored := fmt.Sprintf("\x1b[%smThis is %s\x1b[0m", c.color.Nos(), c.text)
		escapedForm := fmt.Sprintf("%q", colored)
		if scannedLine != escapedForm {
			t.Errorf("Expecting %s, got '%s'\n", escapedForm, scannedLine)
		}
	}
}

func TestContextualSprintf(t *testing.T) {
	testCases := []struct {
		cfmt  SprintfFunc
		color aurora.Color
		text  string
	}{
		{cfmt: Ssuccessf, color: aurora.Color(aurora.GreenFg), text: "success"},
		{cfmt: Sinfof, color: aurora.Color(aurora.CyanFg), text: "info"},
		{cfmt: Swarningf, color: aurora.Color(aurora.BrownFg), text: "warning"},
		{cfmt: Serrorf, color: aurora.Color(aurora.RedFg), text: "error"},
	}

	for _, c := range testCases {
		line := c.cfmt("This is %s", c.text)
		scannedLine := fmt.Sprintf("%q", line)
		colored := fmt.Sprintf("\x1b[%smThis is %s\x1b[0m", c.color.Nos(), c.text)
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
		cfmt  FprintFunc
		color aurora.Color
		text  string
	}{
		{cfmt: Fsuccess, color: aurora.Color(aurora.GreenFg), text: "success"},
		{cfmt: Finfo, color: aurora.Color(aurora.CyanFg), text: "info"},
		{cfmt: Fwarning, color: aurora.Color(aurora.BrownFg), text: "warning"},
		{cfmt: Ferror, color: aurora.Color(aurora.RedFg), text: "error"},
	}

	for _, c := range testCases {
		c.cfmt(rb, c.text)
		line, _ := rb.ReadString('\n')
		scannedLine := fmt.Sprintf("%q", line)
		colored := fmt.Sprintf("\x1b[%sm%s\x1b[0m", c.color.Nos(), c.text)
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
		cfmt  PrintFunc
		color aurora.Color
		text  string
	}{
		{cfmt: Success, color: aurora.Color(aurora.GreenFg), text: "success"},
		{cfmt: Info, color: aurora.Color(aurora.CyanFg), text: "info"},
		{cfmt: Warning, color: aurora.Color(aurora.BrownFg), text: "warning"},
		{cfmt: Error, color: aurora.Color(aurora.RedFg), text: "error"},
	}

	for _, c := range testCases {
		c.cfmt(c.text)
		line, _ := rb.ReadString('\n')
		scannedLine := fmt.Sprintf("%q", line)
		colored := fmt.Sprintf("\x1b[%sm%s\x1b[0m", c.color.Nos(), c.text)
		escapedForm := fmt.Sprintf("%q", colored)
		if scannedLine != escapedForm {
			t.Errorf("Expecting %s, got '%s'\n", escapedForm, scannedLine)
		}
	}
}

func TestContextualSprint(t *testing.T) {
	testCases := []struct {
		cfmt  SprintFunc
		color aurora.Color
		text  string
	}{
		{cfmt: Ssuccess, color: aurora.Color(aurora.GreenFg), text: "success"},
		{cfmt: Sinfo, color: aurora.Color(aurora.CyanFg), text: "info"},
		{cfmt: Swarning, color: aurora.Color(aurora.BrownFg), text: "warning"},
		{cfmt: Serror, color: aurora.Color(aurora.RedFg), text: "error"},
	}

	for _, c := range testCases {
		line := c.cfmt(c.text)
		scannedLine := fmt.Sprintf("%q", line)
		colored := fmt.Sprintf("\x1b[%sm%s\x1b[0m", c.color.Nos(), c.text)
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
		cfmt  FprintlnFunc
		color aurora.Color
		text  string
	}{
		{cfmt: Fsuccessln, color: aurora.Color(aurora.GreenFg), text: "success"},
		{cfmt: Finfoln, color: aurora.Color(aurora.CyanFg), text: "info"},
		{cfmt: Fwarningln, color: aurora.Color(aurora.BrownFg), text: "warning"},
		{cfmt: Ferrorln, color: aurora.Color(aurora.RedFg), text: "error"},
	}

	for _, c := range testCases {
		c.cfmt(rb, c.text)
		line, _ := rb.ReadString(' ')
		scannedLine := fmt.Sprintf("%q", line)
		colored := fmt.Sprintf("\x1b[%sm%s\n\x1b[0m", c.color.Nos(), c.text)
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
		cfmt  PrintlnFunc
		color aurora.Color
		text  string
	}{
		{cfmt: Successln, color: aurora.Color(aurora.GreenFg), text: "success"},
		{cfmt: Infoln, color: aurora.Color(aurora.CyanFg), text: "info"},
		{cfmt: Warningln, color: aurora.Color(aurora.BrownFg), text: "warning"},
		{cfmt: Errorln, color: aurora.Color(aurora.RedFg), text: "error"},
	}

	for _, c := range testCases {
		c.cfmt(c.text)
		line, _ := rb.ReadString(' ')
		scannedLine := fmt.Sprintf("%q", line)
		colored := fmt.Sprintf("\x1b[%sm%s\n\x1b[0m", c.color.Nos(), c.text)
		escapedForm := fmt.Sprintf("%q", colored)
		if scannedLine != escapedForm {
			t.Errorf("Expecting %s, got '%s'\n", escapedForm, scannedLine)
		}
	}
}

func TestContextualSprintln(t *testing.T) {
	testCases := []struct {
		cfmt  SprintlnFunc
		color aurora.Color
		text  string
	}{
		{cfmt: Ssuccessln, color: aurora.Color(aurora.GreenFg), text: "success"},
		{cfmt: Sinfoln, color: aurora.Color(aurora.CyanFg), text: "info"},
		{cfmt: Swarningln, color: aurora.Color(aurora.BrownFg), text: "warning"},
		{cfmt: Serrorln, color: aurora.Color(aurora.RedFg), text: "error"},
	}

	for _, c := range testCases {
		line := c.cfmt(c.text)
		scannedLine := fmt.Sprintf("%q", line)
		colored := fmt.Sprintf("\x1b[%sm%s\n\x1b[0m", c.color.Nos(), c.text)
		escapedForm := fmt.Sprintf("%q", colored)
		if scannedLine != escapedForm {
			t.Errorf("Expecting %s, got '%s'\n", escapedForm, scannedLine)
		}
	}
}
