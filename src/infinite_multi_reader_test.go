package infinite_multi_reader

import (
	"bytes"
	"io"
	"io/ioutil"
	"testing"
)

func TestInfiniteMultiReader(t *testing.T) {
	readers := make(chan io.Reader)
	go func() {
		readers <- bytes.NewBufferString("hello")
		readers <- bytes.NewBufferString(" ")
		readers <- bytes.NewBufferString("world")
		close(readers)
	}()
	r := InfiniteMultiReader(func() (reader io.Reader, err error) {
		reader, ok := <-readers
		if !ok {
			err = io.EOF
		}
		return
	})
	result, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	if string(result) != "hello world" {
		t.Fatal("Unexpected result:", string(result))
	}
}
