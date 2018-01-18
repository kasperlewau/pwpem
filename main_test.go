package pwpem_test

import (
	"bytes"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/kasperlewau/pwpem"
)

func TestEncryptDecrypt(t *testing.T) {
	pass := []byte("password")
	fixtures, err := filepath.Glob("fixtures/*")
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range fixtures {
		t.Run(f, func(t *testing.T) {
			b, err := ioutil.ReadFile(f)
			if err != nil {
				t.Fatal(err)
			}
			e, err := pwpem.Encrypt(b, pass)
			if err != nil {
				blk, _ := pem.Decode(b)
				if x509.IsEncryptedPEMBlock(blk) {
					if err != pwpem.ErrBlockEncrypted {
						t.Fatal("Failed with bad err", err)
					}
					return
				}
				t.Fatal(err)
			}
			d, err := pwpem.Decrypt(e, pass)
			if err != nil {
				t.Fatal(err)
			}
			if !bytes.Equal(b, d) {
				t.Fatalf("Expected encrypt(cert, pw) -> drypt(cert, pw) to equal original %v. Got %v", b, d)
			}
		})
	}
}

func BenchmarkEncrypt(b *testing.B) {
	b.ReportAllocs()
	c := []byte("-----BEGIN EC PRIVATE KEY-----\nMHcCAQEEIG7rIirIoNGikSvdKGQPzpmSzHDHuzgbyKanuMb1ew0coAoGCCqGSM49\nAwEHoUQDQgAEr6e6UTc8GM1cFff3VbTiT/NxH/Ebc4fyV106mKH9S7YxCWxe+sCz\nlWzcrox9ONM5bJVx0aXG2ce8pQzHlcbVJA==\n-----END EC PRIVATE KEY-----\n")
	p := []byte("pass")
	for i := 0; i < b.N; i++ {
		pwpem.Encrypt(c, p)
	}
}

func BenchmarkDecrypt(b *testing.B) {
	b.ReportAllocs()
	c := []byte("-----BEGIN EC PRIVATE KEY-----\nMHcCAQEEIG7rIirIoNGikSvdKGQPzpmSzHDHuzgbyKanuMb1ew0coAoGCCqGSM49\nAwEHoUQDQgAEr6e6UTc8GM1cFff3VbTiT/NxH/Ebc4fyV106mKH9S7YxCWxe+sCz\nlWzcrox9ONM5bJVx0aXG2ce8pQzHlcbVJA==\n-----END EC PRIVATE KEY-----\n")
	p := []byte("pass")
	e, err := pwpem.Encrypt(c, p)
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		pwpem.Decrypt(e, p)
	}
}
