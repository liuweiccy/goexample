package sorting

import (
	"fmt"
	"sort"
	"testing"
)

func TestPrintTracks(t *testing.T) {
	PrintTracks(Tracks)

	fmt.Println()

	sort.Sort(ByArtist(Tracks))
	PrintTracks(Tracks)

	sort.Reverse(ByArtist(Tracks))
	PrintTracks(Tracks)

	fmt.Errorf()
}
