# GzipCloser
Gzip Closer is a wrapper around golang [gzip package](https://golang.org/pkg/compress/gzip/) where the underlying gzip stream is also closed on Close(). Useful if when you don't want to keep track of 2 readers/writers.


[![Go Report Card](https://goreportcard.com/badge/github.com/zatte/gzipcloser?style=flat-square)](https://goreportcard.com/report/github.com/zatte/gzipcloser)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/zatte/gzipcloser)

## License
Apache 2.0

## Usage

```golang

  // Reading
  f, _ := os.Open("somefile.gz")
  gzReader := gzipcloser.NewReader(f)
  gzReader.Read(...)
  gzReader.Close() // same as gzReader.Close() && f.Close()


  // Writing
  f, _ = os.OpenFile("somefile.gz", os.O_RDWR|os.O_CREATE, 0660)
  gzWriter := gzipcloser.NewWriter(f)
  gzWriter.Write(...)
  gzWriter.Close() // same as gzWriter.Flush() && gzWriter.Close() && f.Flush() && f.Close(); f.Flush is only called if f implements Flush()


  // Writing to a writer without Flush() and Close()
  b, _ := bytes.NewBuffer(nil)
  gzWriter := gzipcloser.NewWriter(b)
  gzWriter.Write(...)
  gzWriter.Close() // same as gzWriter.Flush() && gzWriter.Close()

```

`Flush()` and `Close()` calla are only invoked on the base reader/writer if if implements the methods.