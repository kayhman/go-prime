package main

import (
	"fmt"
	"flag"
)


func main() {
	flag.Parse()
	max := *maxPtr
	
	public, private, err := generateKeys(max)
	if err != nil {
		fmt.Println("Failed to generate keys")
		return
	}

	c, _ := encrypt(public, *msgPtr)
	m, _ := decrypt(private, c)

	fmt.Printf("chiffrement: %s \n", c.String())
	fmt.Printf("decriffrement: %s \n", m.String())

	return
	
	pu, err := loadPublicKey("pub2")
	fmt.Printf("key: %s \n", pu)

	pv, err := loadPrivateKey("key.pub")
	fmt.Printf("key: %s \n", pv)

	c, _ = encrypt(pu, *msgPtr)
	m, _ = decrypt(pv, c)

	fmt.Printf("crypto: %s \n", c.String())
	fmt.Printf("decrypto: %s \n", m.String())
}
