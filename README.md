# InfiniteMultiReader

Implements a Reader which can read from a (potentially) infinite sequence of other readers. Use like:

```Go
getNextReader := func() (io.Reader, error) {
  // ...
  // return the_next_reader, err
}
reader := InfiniteMultiReader(getNextReader)
```

The ```getNextReader``` function should ```return nil, io.EOF``` when there are no more readers to consume.
