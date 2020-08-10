package stdout

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var spinChars = `|/-\`

type Spinner struct {
	message string
	i       int
}

func NewSpinner(message string) *Spinner {
	return &Spinner{message: message}
}

func (s *Spinner) Tick() {
	fmt.Printf("%c %s \r", spinChars[s.i], s.message)
	s.i = (s.i + 1) % len(spinChars)
}

func isTTY() bool {
	fi, err := os.Stdout.Stat()
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeCharDevice != 0
}

//ConsoleTreat blah blah
func ConsoleTreat(logmsg string) {
	flag.Parse()
	s := NewSpinner(logmsg)
	isTTY := isTTY()

	if isTTY {
		s.Tick()
	}
	time.Sleep(100 * time.Millisecond)

}
