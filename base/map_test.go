package base

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"testing"
)

func TestMap1(t *testing.T) {
	ages := make(map[string]int, 10)

	ages["eric"] = 29
	ages["chris"] = 2
	ages["chunyan"] = 28
	fmt.Println(ages)

	delete(ages, "dd")
	fmt.Println(ages)

	delete(ages, "eric")
	fmt.Println(ages)

	m1 := &sync.Map{}
	m1.Store("Liu", 23)
	m1.Store("Chen", 20)
	m1.Range(func(key, value interface{}) bool {
		k, v := key.(string), value.(int)
		fmt.Println(k, "===", v)
		return true
	})
}

func TestMap2(t *testing.T) {
	seen := make(map[string]bool)

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup:%v\n", err)
		os.Exit(1)
	}
}
