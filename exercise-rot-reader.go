package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	c := make([]byte, len(b))
	n, e := r.r.Read(c)
	if e != nil {
		return n, e
	}
	for i := 0; i < n; i++ {
		character := c[i]
		var code byte
		switch {
		case 'A' <= character && character <= 'Z':
			code = character + 13
			if code > 'Z' {
				code = 'A' - 1 + (code - 'Z')
			}
		case 'a' <= character && character <= 'z':
			code = character + 13
			if code > 'z' {
				code = 'a' - 1 + (code - 'z')
			}
		default:
			code = character
		}
		b[i] = code
	}
	return len(c), nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbgr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
