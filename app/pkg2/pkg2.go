package pkg2

import (
    "bytes"
    "compress/gzip"
)

func Work() string {
    var buf bytes.Buffer
    gz := gzip.NewWriter(&buf)
    for i := 0; i < 10000; i++ {
        gz.Write([]byte("hello world"))
    }
    gz.Close()
    return "pkg2 "
}
