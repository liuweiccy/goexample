package strings

import (
	"io"
)

type StringRead struct {
	b []byte
}

func (s *StringRead) Read(p []byte) (n int, err error) {
	n = copy(p, s.b[:])
	return
}

func NewReader(s string) io.Reader {
	sr := &StringRead{b: []byte(s)}
	return sr
}

type limitRead struct {
	r     io.Reader
	n     int64
	count int64 // 已经读取的数据量
}

func (l *limitRead) Read(p []byte) (n int, err error) {
	if l.count >= l.n {
		return 0, io.EOF
	}
	n, err = l.r.Read(p)
	if err != nil {
		return n, err
	}
	l.count += int64(n)
	if l.count >= l.n {
		err = io.EOF
	}
	return n, err
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitRead{r: r, n: n}
}
