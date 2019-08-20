package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	// Unicode字符数量
	count := make(map[rune]int)
	// UTF-8编码长度
	var utfLen [utf8.UTFMax + 1]int
	// 非法UTF-8字符数量
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}

		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "charcount:%v\n", err)
			os.Exit(1)
		}

		fmt.Println("输入了字符：", r)

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		count[r]++
		utfLen[n]++
	}
}
