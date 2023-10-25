// ex7.5 provides a LimitReader that reports EOF at a given offset.
package reader

import "io"

type limitReader struct {
	r        io.Reader
	n, limit int
}

func (l *limitReader) Read(p []byte) (n int, err error) {
	n, err = l.r.Read(p[:l.limit])
	l.n += n
	if l.n >= l.limit {
		err = io.EOF
	}
	return
}

func LimitReader(r io.Reader, limit int) io.Reader {
	return &limitReader{r: r, limit: limit}
}
