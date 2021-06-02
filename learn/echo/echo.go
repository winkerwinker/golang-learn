package echo

import (
	"flag"
	"os"
)

var omitNewline = flag.Bool("n", false, "don't print final newline")

const (
	Space   = " "
	Newline = "\n"
)

// todo go run 如何穿参数
func Echo() {
	flag.Parse() // Scans the arg list and sets up flags
	var s string = ""
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += Space
		}
		s += flag.Arg(i)
	}
	if !*omitNewline {
		s += Newline
	}
	// do not handler
	os.Stdout.WriteString(s)
}
