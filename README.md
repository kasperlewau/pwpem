# pwpem [![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/github.com/kasperlewau/pwpem)
> encrypt & decrypt PEM block(s) with passwords


# install
```sh
go get github.com/kasperlewau/pempw
```

# usage
```go
import "github.com/kasperlewau/pempw"

pem := []byte("-----BEGIN PRIVATE KEY-----\nbase64encodedprivatekey....")
pass := []byte("password")

e, err := pwpem.Encrypt(pem, pass)
fmt.Println(string(e)); // "-----BEGIN PRIVATE KEY-----\nProc-Type: 4,ENCRYPTED\nDEK-Info: AES-256....."

d, err := pwpem.Decrypt(e, pass)
fmt.Println(string(d)); // "-----BEGIN PRIVATE KEY-----\nbase64encodedprivatekey....")
```
