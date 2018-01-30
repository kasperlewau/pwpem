# pwpem [![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/github.com/kasperlewau/pwpem) [![Build Status](https://travis-ci.org/kasperlewau/pwpem.svg?branch=master)](https://travis-ci.org/kasperlewau/pwpem)
> encrypt & decrypt PEM block(s) with passwords


# install
```sh
go get github.com/kasperlewau/pwpem
```

# usage
```go
import "github.com/kasperlewau/pwpem"

pem := []byte("-----BEGIN PRIVATE KEY-----\nbase64encodedprivatekey....")
pass := []byte("password")

e, err := pwpem.Encrypt(pem, pass)
fmt.Println(string(e)); // "-----BEGIN PRIVATE KEY-----\nProc-Type: 4,ENCRYPTED\nDEK-Info: AES-256....."

d, err := pwpem.Decrypt(e, pass)
fmt.Println(string(d)); // "-----BEGIN PRIVATE KEY-----\nbase64encodedprivatekey....")
```
