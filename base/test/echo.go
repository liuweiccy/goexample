package test

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	n = flag.Bool("n", false, "换行符")
	s = flag.String("s", " ", "分割符")
)

var out io.Writer = os.Stdout

func echo(newline bool, sep string, args []string) error {
	_, _ = fmt.Fprint(out, strings.Join(args, sep))
	if newline {
		_, _ = fmt.Fprintln(out)
	}
	return nil
}
