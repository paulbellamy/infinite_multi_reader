package infinite_multi_reader

import (
	"bytes"
	"io"
)

type infiniteMultiReader struct {
	currentReader io.Reader
	nextReader    func() (io.Reader, error)
}

func (imr *infiniteMultiReader) Read(p []byte) (n int, err error) {
	for {
		n, err = imr.currentReader.Read(p)
		if n > 0 || err != io.EOF {
			if err == io.EOF {
				// Don't return EOF yet. There may be more bytes
				// in the remaining readers.
				err = nil
			}
			return
		}
		if imr.currentReader, err = imr.nextReader(); err != nil {
			return
		}
	}
	return 0, io.EOF
}

// InfiniteMultiReader returns a Reader that's the logical concatenation of the
// input readers returned by the nextReader function. They're read
// sequentially. When there are no more readers, the nextReader function should
// return the io.EOF error. Once all inputs are drained, Read will return
// io.EOF.
func InfiniteMultiReader(nextReader func() (io.Reader, error)) io.Reader {
	return &infiniteMultiReader{
		currentReader: &bytes.Buffer{},
		nextReader:    nextReader,
	}
}
