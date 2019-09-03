package tempconv

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"testing"
)

var temp = CelsiusFlag("temp", 20.0, "the temperature")

func TestCelsiusFlag(t *testing.T) {
	flag.Parse()
	fmt.Println(*temp)
}

func TestType(t *testing.T) {
	var w io.Writer
	fmt.Printf("%T\n", w)

	w = os.Stdout
	fmt.Printf("%T\n", w)

	w = new(bytes.Buffer)
	fmt.Printf("%T\n", w)
}

const debug = false

func TestNil1(t *testing.T) {
	var buf io.Writer
	fmt.Println(buf)

	if debug {
		buf = new(bytes.Buffer)
	}

	f(buf)
}

func f(out io.Writer) {
	fmt.Println(out)
	fmt.Println(out == nil)

	if out != nil {
		out.Write([]byte("done\n"))
	}
}
