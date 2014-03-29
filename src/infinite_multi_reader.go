package infinite_multi_reader

import (
	"bytes"
	"io"
)

var EmptyReader = &bytes.Buffer{}

type infiniteMultiReader struct {
	currentReader io.Reader
	callback      func() (io.Reader, error)
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
		if imr.currentReader, err = imr.callback(); err != nil {
			return
		}
	}
	return 0, io.EOF
}

// MultiReader returns a Reader that's the logical concatenation of
// the provided input readers.  They're read sequentially.  Once all
// inputs are drained, Read will return EOF.
func InfiniteMultiReader(callback func() (io.Reader, error)) io.Reader {
	return &infiniteMultiReader{
		currentReader: EmptyReader,
		callback:      callback,
	}
}
