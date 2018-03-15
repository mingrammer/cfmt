package cfmt

import (
	"io"

	"github.com/fatih/color"
)

// Fsuccessf writes green colored text in manner of fmt.Fprintf
func Fsuccessf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	n, err = color.New(color.FgGreen).Fprintf(w, format, a...)
	return
}

// Successf prints green colored text in manner of fmt.Printf
func Successf(format string, a ...interface{}) (n int, err error) {
	n, err = Fsuccessf(color.Output, format, a...)
	return
}

// Ssuccessf returns green colored string in manner of fmt.Sprintf
func Ssuccessf(format string, a ...interface{}) string {
	return color.New(color.FgGreen).Sprintf(format, a...)
}

// Finfof writes cyan colored text in manner of fmt.Fprintf
func Finfof(w io.Writer, format string, a ...interface{}) (n int, err error) {
	n, err = color.New(color.FgCyan).Fprintf(w, format, a...)
	return
}

// Infof prints cyan colored text in manner of fmt.Printf
func Infof(format string, a ...interface{}) (n int, err error) {
	n, err = Finfof(color.Output, format, a...)
	return
}

// Sinfof returns cyan colored string in manner of fmt.Sprintf
func Sinfof(format string, a ...interface{}) string {
	return color.New(color.FgCyan).Sprintf(format, a...)
}

// Fwarningf writes yellow colored text in manner of fmt.Fprintf
func Fwarningf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	n, err = color.New(color.FgYellow).Fprintf(w, format, a...)
	return
}

// Warningf prints yellow colored text in manner of fmt.Printf
func Warningf(format string, a ...interface{}) (n int, err error) {
	n, err = Fwarningf(color.Output, format, a...)
	return
}

// Swarningf returns yellow colored string in manner of fmt.Sprintf
func Swarningf(format string, a ...interface{}) string {
	return color.New(color.FgYellow).Sprintf(format, a...)
}

// Ferrorf writes red colored text in manner of fmt.Fprintf
func Ferrorf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	n, err = color.New(color.FgRed).Fprintf(w, format, a...)
	return
}

// Errorf prints red colored text in manner of fmt.Printf
func Errorf(format string, a ...interface{}) (n int, err error) {
	n, err = Ferrorf(color.Output, format, a...)
	return
}

// Serrorf returns red colored string in manner of fmt.Sprintf
func Serrorf(format string, a ...interface{}) string {
	return color.New(color.FgRed).Sprintf(format, a...)
}

// Fsuccess prints green colored text in manner of fmt.Fprint
func Fsuccess(w io.Writer, a ...interface{}) (n int, err error) {
	n, err = color.New(color.FgGreen).Fprint(w, a...)
	return
}

// Success prints green colored text in manner of fmt.Print
func Success(a ...interface{}) (n int, err error) {
	n, err = Fsuccess(color.Output, a...)
	return
}

// Ssuccess returns green colored string in manner of fmt.Sprint
func Ssuccess(a ...interface{}) string {
	return color.New(color.FgGreen).Sprint(a...)
}

// Finfo prints cyan colored text in manner of fmt.Fprint
func Finfo(w io.Writer, a ...interface{}) (n int, err error) {
	n, err = color.New(color.FgCyan).Fprint(w, a...)
	return
}

// Info prints cyan colored text in manner of fmt.Print
func Info(a ...interface{}) (n int, err error) {
	n, err = Finfo(color.Output, a...)
	return
}

// Sinfo returns cyan colored string in manner of fmt.Sprint
func Sinfo(a ...interface{}) string {
	return color.New(color.FgCyan).Sprint(a...)
}

// Fwarning prints yellow colored text in manner of fmt.Fprint
func Fwarning(w io.Writer, a ...interface{}) (n int, err error) {
	n, err = color.New(color.FgYellow).Fprint(w, a...)
	return
}

// Warning prints yellow colored text in manner of fmt.Print
func Warning(a ...interface{}) (n int, err error) {
	n, err = Fwarning(color.Output, a...)
	return
}

// Swarning returns yellow colored string in manner of fmt.Sprint
func Swarning(a ...interface{}) string {
	return color.New(color.FgYellow).Sprint(a...)
}

// Ferror prints red colored text in manner of fmt.Fprint
func Ferror(w io.Writer, a ...interface{}) (n int, err error) {
	n, err = color.New(color.FgRed).Fprint(w, a...)
	return
}

// Error prints red colored text in manner of fmt.Print
func Error(a ...interface{}) (n int, err error) {
	n, err = Ferror(color.Output, a...)
	return
}

// Serror returns red colored string in manner of fmt.Sprint
func Serror(a ...interface{}) string {
	return color.New(color.FgRed).Sprint(a...)
}

// Fsuccessln prints green colored text in manner of fmt.Fprintln
func Fsuccessln(w io.Writer, a ...interface{}) (n int, err error) {
	n, err = color.New(color.FgGreen).Fprintln(w, a...)
	return
}

// Successln prints green colored text in manner of fmt.Println
func Successln(a ...interface{}) (n int, err error) {
	n, err = Fsuccessln(color.Output, a...)
	return
}

// Ssuccessln returns green colored string in manner of fmt.Sprintln
func Ssuccessln(a ...interface{}) string {
	return color.New(color.FgGreen).Sprintln(a...)
}

// Finfoln prints cyan colored text in manner of fmt.Fprintln
func Finfoln(w io.Writer, a ...interface{}) (n int, err error) {
	n, err = color.New(color.FgCyan).Fprintln(w, a...)
	return
}

// Infoln prints cyan colored text in manner of fmt.Println
func Infoln(a ...interface{}) (n int, err error) {
	n, err = Finfoln(color.Output, a...)
	return
}

// Sinfoln returns cyan colored string in manner of fmt.Sprintln
func Sinfoln(a ...interface{}) string {
	return color.New(color.FgCyan).Sprintln(a...)
}

// Fwarningln prints yellow colored text in manner of fmt.Fprintln
func Fwarningln(w io.Writer, a ...interface{}) (n int, err error) {
	n, err = color.New(color.FgYellow).Fprintln(w, a...)
	return
}

// Warningln prints yellow colored text in manner of fmt.Println
func Warningln(a ...interface{}) (n int, err error) {
	n, err = Fwarningln(color.Output, a...)
	return
}

// Swarningln returns yellow colored string in manner of fmt.Sprintln
func Swarningln(a ...interface{}) string {
	return color.New(color.FgYellow).Sprintln(a...)
}

// Ferrorln prints red colored text in manner of fmt.Fprintln
func Ferrorln(w io.Writer, a ...interface{}) (n int, err error) {
	n, err = color.New(color.FgRed).Fprintln(w, a...)
	return
}

// Errorln prints red colored text in manner of fmt.Println
func Errorln(a ...interface{}) (n int, err error) {
	n, err = Ferrorln(color.Output, a...)
	return
}

// Serrorln returns red colored string in manner of fmt.Sprintln
func Serrorln(a ...interface{}) string {
	return color.New(color.FgRed).Sprintln(a...)
}
