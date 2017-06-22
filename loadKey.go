package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

func loadPublicKey(file string) (BigKey, error) {
	pubPEM, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
		return BigKey{}, err
	}

	block, _ := pem.Decode([]byte(pubPEM))
	if block == nil {
		panic("failed to parse PEM block containing the public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic("failed to parse DER encoded public key: " + err.Error())
	}

	var key BigKey
	switch pb := pub.(type) {
	case *rsa.PublicKey:
		key.modulus = *pb.N
		key.exponent.SetInt64(int64(pb.E))
	default:
		panic("unknown type of public key")
	}
	return key, nil
}

func loadPrivateKey(file string) (BigKey, error) {
	pubPEM, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
		return BigKey{}, err
	}

	block, _ := pem.Decode([]byte(pubPEM))
	if block == nil {
		panic("failed to parse PEM block containing the public key")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic("failed to parse DER encoded public key: " + err.Error())
	}

	var key BigKey
	key.modulus = *priv.N
	key.exponent = *priv.D

	for i, p := range priv.Primes {
		fmt.Printf("prime %d : %s", i, p.String())
	}

	return key, nil
}
