package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"encoding/pem"
	"os"
)

func main() {

	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	
	data, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		panic(err)
	}
	fmt.Println(hex.EncodeToString(data))
	
	pemKey := pem.Block{Type: "EC" + " PRIVATE KEY", Bytes: data}
	keyOut, err := os.Create("ctm-private.pem")
	if err != nil {
		panic(err)
	}

	keyOut.Chmod(0600)
	defer keyOut.Close()
	pem.Encode(keyOut, &pemKey)

	// Once the ctm-private.pem is generated, 
	// run: openssl pkey -in ctm-private.pem -pubout -out ctm-public.pem
	// to create the public key
}
