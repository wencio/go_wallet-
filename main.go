package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {

	// Generate a key
	pvk, err := crypto.GenerateKey()
	// Check for errors
	if err != nil {

		log.Fatalf("Error generating key %v", err)
	}
	// Create de private key usin ECDSA - Elliptic Curve Digital Signature Algorithms
	pData := crypto.FromECDSA(pvk)
	fmt.Println("Private key:", hexutil.Encode(pData))
	// Creating public address from private key
	pubData := crypto.FromECDSAPub(&pvk.PublicKey)

	fmt.Println("Public key in bytes array:", pubData)
	fmt.Println("Public Key in Hex:", hexutil.Encode(pubData))
	// Creating Public Address from Public Key
	fmt.Println("Public address in Hex:", crypto.PubkeyToAddress(pvk.PublicKey).Hex())
}
