package pkg1

import (
    "crypto/x509"
    "math/big"
)

func Work() string {
    var x big.Int
    for i := 0; i < 5000; i++ {
        x.Mul(&x, big.NewInt(int64(i+1)))
    }
    _ = x509.Certificate{}
    return "pkg1 "
}
