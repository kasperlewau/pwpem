// Package pwpem encrypts & decrypts PEM blocks with a password
package pwpem // github.com/kasperlewau/pwpem

import (
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

var (
	// ErrNoBlock signals that there was no block to be encrypted/decrypted in the given byte slice
	ErrNoBlock = errors.New("pwpem: no block found")
	// ErrBlockEncrypted signals that the first block is already encrypted
	ErrBlockEncrypted = errors.New("pwpem: first block is already encrypted")
)

// Encrypt encrypts the first encountered PEM block in the byte slice with password.
//
// The full encrypted chain is returned as a byte slice.
func Encrypt(c, p []byte) ([]byte, error) {
	blk, rest := pem.Decode(c)
	if blk == nil {
		return nil, ErrNoBlock
	}
	if x509.IsEncryptedPEMBlock(blk) {
		return nil, ErrBlockEncrypted
	}
	enc, err := x509.EncryptPEMBlock(rand.Reader, blk.Type, blk.Bytes, p, x509.PEMCipherAES256)
	if err != nil {
		return nil, err
	}
	b := pem.EncodeToMemory(enc)
	return append(b, rest...), nil
}

// Decrypt decrypts the first encountered PEM block in the byte slice with password.
//
// The full decrypted chain is returned as a byte slice.
func Decrypt(c, p []byte) ([]byte, error) {
	blk, rest := pem.Decode(c)
	if blk == nil {
		return nil, ErrNoBlock
	}
	der, err := x509.DecryptPEMBlock(blk, p)
	if err != nil {
		return nil, err
	}
	b := pem.EncodeToMemory(&pem.Block{Type: blk.Type, Bytes: der})
	return append(b, rest...), nil
}
