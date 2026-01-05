package pkg1

import (
	"crypto/x509"
	"net/http/httptest"
	"compress/zlib"
	"bytes"
	"io"
)

func Work() string {
	// Force heavy stdlib compilation
	cert := &x509.Certificate{}
	_ = cert

	req := httptest.NewRequest("GET", "/", nil)
	_ = req

	var buf bytes.Buffer
	w := zlib.NewWriter(&buf)
	for i := 0; i < 50000; i++ {
		w.Write([]byte("this forces compression work"))
	}
	w.Close()

	io.Copy(io.Discard, &buf)

	return "pkg1 "
}
