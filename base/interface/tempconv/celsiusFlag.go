package tempconv

import (
	"flag"
	"fmt"
)

type celsiusFlag struct {
	Celsius
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	_, _ = fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "c":
		f.Celsius = Celsius(value)
		return nil
	case "F", "f":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temp %q\n", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}