package sleep

import (
	"flag"
	"fmt"
	"testing"
	"time"
)

var period = flag.Duration("period", time.Second, "sleep duration")

func TestPeriod(t *testing.T)  {
	flag.Parse()
	fmt.Printf("Sleeping for %v ...\n", *period)
	time.Sleep(*period)
}
